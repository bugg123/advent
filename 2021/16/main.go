package main

import (
	"encoding/hex"
	"fmt"
	"math"
)

func main() {
	// input := in
	// fmt.Println(input)

	for str, ans := range tests {
		bytes, err := hex.DecodeString(str)
		if err != nil {
			panic(err)
		}
		b := &BitStream{
			Bytes: bytes,
			Pos:   0,
		}
		for _, b := range bytes {
			fmt.Printf("%08b ", b)
		}
		fmt.Println()
		packet, _ := b.GetPacket()
		actAns := getValue(*packet)
		if ans != actAns {
			fmt.Printf("Failed packet: %v\n", packet)
			fmt.Printf("Got: %d, Wanted %d\n", actAns, ans)
			panic("FAIL")
		}
	}
	bytes, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	b := &BitStream{
		Bytes: bytes,
		Pos:   0,
	}
	for _, b := range bytes {
		fmt.Printf("%08b ", b)
	}
	fmt.Println()
	packet, _ := b.GetPacket()
	fmt.Printf("Top: %+v\n", packet)
	for !b.IsEmpty() {
		fmt.Printf("%b", b.GetBit())
	}

	fmt.Println()
	fmt.Println(getVersionSum([]Packet{*packet}))
	fmt.Println(getValue(*packet))
	// bytes, err := hex.DecodeString(testOperator)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, b := range bytes {
	// 	fmt.Printf("%08b ", b)
	// }
	// fmt.Println()
	// fmt.Println(getOperatorFromPacket(bytes))
}

func getVersionSum(packets []Packet) int {
	sum := 0
	for _, p := range packets {
		sum += p.Version
		sum += getVersionSum(p.SubPackets)
	}
	return sum
}

func getValue(packet Packet) int {
	val := 0
	switch packet.ID {
	case 4:
		return packet.Val
	case 0:
		for _, subPacket := range packet.SubPackets {
			val += getValue(subPacket)
		}
	case 1:
		val = 1
		for _, subPacket := range packet.SubPackets {
			val *= getValue(subPacket)
		}
	case 2:
		val = math.MaxInt
		for _, subPacket := range packet.SubPackets {
			val = getMin(val, getValue(subPacket))
		}
	case 3:
		val = math.MinInt
		for _, subPacket := range packet.SubPackets {
			val = getMax(val, getValue(subPacket))
		}
	case 5:
		if len(packet.SubPackets) != 2 {
			panic("Invalid greater than packet")
		}
		if getValue(packet.SubPackets[0]) > getValue(packet.SubPackets[1]) {
			val = 1
		}
	case 6:
		if len(packet.SubPackets) != 2 {
			fmt.Printf("Length: %d, %v\n", len(packet.SubPackets), packet)
			panic("Invalid less than packet")
		}
		if getValue(packet.SubPackets[0]) < getValue(packet.SubPackets[1]) {
			val = 1
		}
	case 7:
		if len(packet.SubPackets) != 2 {
			panic("Invalid equal to packet")
		}
		fmt.Printf("Values: %d %d\n", getValue(packet.SubPackets[0]), getValue(packet.SubPackets[1]))
		if getValue(packet.SubPackets[0]) == getValue(packet.SubPackets[1]) {
			val = 1
		}

	}
	return val
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (b *BitStream) GetPacket() (*Packet, int) {
	consumed := 6
	packet := b.GetNextPacket()
	if packet == nil {
		return nil, 6
	}
	var cons int
	switch packet.ID {
	case 4:
		val, cons := b.GetLiteral()
		packet.Val = val
		consumed += cons
	default:
		packet.SubPackets, cons = b.GetOperatorPackets()
	}
	return packet, consumed + cons
}

func (b *BitStream) GetOperatorPackets() ([]Packet, int) {
	lengthType := b.GetBit()
	consumed := 1
	var packets []Packet
	var cons int
	if lengthType == 0 {
		packets,cons = b.GetZeroOperator()
	} else if lengthType == 1 {
		packets, cons = b.GetOneOperator()
	} else {
		panic("Unknown operator type")
	}
	return packets, consumed + cons
}

func (b *BitStream) GetOneOperator() ([]Packet, int) {
	fmt.Println("One operator")
	length := b.GetBitsAsInt(11)
	consumed := 11
	fmt.Printf("Length: %d\n", length)
	subPackets := make([]Packet, 0)
	for length > 0 {
		pack, cons := b.GetPacket()
		consumed += cons
		if pack == nil || EmptyPacket(*pack) {
			return subPackets, consumed
		}
		subPackets = append(subPackets, *pack)
		length--
	}
	if length < 0 {
		panic("Negative length")
	}
	return subPackets, consumed
}

func (b *BitStream) GetZeroOperator() ([]Packet, int) {
	fmt.Println("Zero operator")
	length := b.GetBitsAsInt(15)
	consumed := 15
	fmt.Printf("Length: %d\n", length)
	subPackets := make([]Packet, 0)
	for length > 0 {
		// b.GetNextPacket()
		// val, consumed := b.GetLiteral()
		pack, cons := b.GetPacket()
		consumed += cons
		if pack == nil || EmptyPacket(*pack) {
			return subPackets, consumed
		}
		subPackets = append(subPackets, *pack)
		length -= cons
	}
	if length < 0 {
		panic("Negative length")
	}

	// b.NextByte()
	return subPackets, consumed
}

func EmptyPacket(p Packet) bool {
	return p.ID == 0 && len(p.SubPackets) == 0 && p.Val == 0 && p.Version == 0
}

func (b *BitStream) GetLiteral() (int, int) {
	// fmt.Println("Getting Literal")
	keepGoing := true
	res := 0
	bytesConsumed := 0
	for keepGoing {
		bytesConsumed++
		if b.GetBit() == 0 {
			keepGoing = false
		}
		for i := 0; i < 4; i++ {
			res = res << 1
			res = res | int(b.GetBit())
			bytesConsumed++
		}
	}
	// fmt.Println(res)
	// bytesConsumed += b.NextByte()
	return res, bytesConsumed
}

type Packet struct {
	Version    int
	ID         int
	Val        int
	SubPackets []Packet
}

type BitStream struct {
	Bytes []byte
	Pos   int
}

func (b *BitStream) GetBit() byte {
	if b.IsEmpty() {
		return 0
	}
	curByte := b.Bytes[0]
	bit := curByte << b.Pos >> 7
	b.Pos++
	if b.Pos > 7 {
		b.Bytes = b.Bytes[1:]
		b.Pos = 0
	}
	return bit
}

func (b *BitStream) IsEmpty() bool {
	return len(b.Bytes) == 0
}

func (b *BitStream) GetBitsAsInt(n int) int {
	res := 0
	for i := 0; i < n; i++ {
		res = res << 1
		res = res | int(b.GetBit())
	}
	return res
}

func (b *BitStream) GetNextPacket() *Packet {
	if b.IsEmpty() {
		return nil
	}
	packet := &Packet{
		Version: int(b.GetBitsAsInt(3)),
		ID:      int(b.GetBitsAsInt(3)),
	}

	return packet
}

type Bit bool

func (b Bit) String() string {
	if b {
		return "1"
	}
	return "0"
}

var testInput = "D2FE28"
var testOperator = "38006F45291200"
var testOperator2 = "EE00D40C823060"
var testOperator3 = "8A004A801A8002F478"
var testOperator4 = "620080001611562C8802118E34"
var testOperator5 = "C0015000016115A2E0802F182340"
var testOperator6 = "A0016C880162017C3686B18A3D4780"

var tests = map[string]int{
	"C200B40A82":                 3,
	"04005AC33890":               54,
	"880086C3E88112":             7,
	"CE00C43D881120":             9,
	"D8005AC2A8F0":               1,
	"F600BC2D8F":                 0,
	"9C005AC2F8F0":               0,
	"9C0141080250320F1802104A08": 1,
}

var input = "220D62004EF14266BBC5AB7A824C9C1802B360760094CE7601339D8347E20020264D0804CA95C33E006EA00085C678F31B80010B88319E1A1802D8010D4BC268927FF5EFE7B9C94D0C80281A00552549A7F12239C0892A04C99E1803D280F3819284A801B4CCDDAE6754FC6A7D2F89538510265A3097BDF0530057401394AEA2E33EC127EC3010060529A18B00467B7ABEE992B8DD2BA8D292537006276376799BCFBA4793CFF379D75CA1AA001B11DE6428402693BEBF3CC94A314A73B084A21739B98000010338D0A004CF4DCA4DEC80488F004C0010A83D1D2278803D1722F45F94F9F98029371ED7CFDE0084953B0AD7C633D2FF070C013B004663DA857C4523384F9F5F9495C280050B300660DC3B87040084C2088311C8010C84F1621F080513AC910676A651664698DF62EA401934B0E6003E3396B5BBCCC9921C18034200FC608E9094401C8891A234080330EE31C643004380296998F2DECA6CCC796F65224B5EBBD0003EF3D05A92CE6B1B2B18023E00BCABB4DA84BCC0480302D0056465612919584662F46F3004B401600042E1044D89C200CC4E8B916610B80252B6C2FCCE608860144E99CD244F3C44C983820040E59E654FA6A59A8498025234A471ED629B31D004A4792B54767EBDCD2272A014CC525D21835279FAD49934EDD45802F294ECDAE4BB586207D2C510C8802AC958DA84B400804E314E31080352AA938F13F24E9A8089804B24B53C872E0D24A92D7E0E2019C68061A901706A00720148C404CA08018A0051801000399B00D02A004000A8C402482801E200530058AC010BA8018C00694D4FA2640243CEA7D8028000844648D91A4001088950462BC2E600216607480522B00540010C84914E1E0002111F21143B9BFD6D9513005A4F9FC60AB40109CBB34E5D89C02C82F34413D59EA57279A42958B51006A13E8F60094EF81E66D0E737AE08"
