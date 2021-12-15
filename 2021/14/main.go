package main

import (
	"bytes"
	"fmt"
	"math"
)

func main() {

	pairs := testPairs
	template := []byte(testTemplate)
	count := make(map[string]map[string]int64, 0)

	for str := range pairs {
		count[str+"-0"] = countInitial(str)
	}

	for str := range pairs {
		getCount(str, 40, pairs, count)
	}

	// for i := 0; i < 41; i++ {
	// for k, v := range count {
	// 	if strings.HasSuffix(k, fmt.Sprintf("-%d", i)) {
	// 		fmt.Printf("%v: %v\n", k, v)
	// 	}
	// }
	// }

	// for str, count := range count {
	// 	fmt.Printf("%v: %v\n", str, count)
	// }
	fmt.Printf("NN-40 Pairs: %v\n", count["NN-40"])
	fmt.Printf("NC-40 Pairs: %v\n", count["NC-40"])
	fmt.Printf("CB-40 Pairs: %v\n", count["CB-40"])

	finalCount := make(map[string]int)
	pairCount := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		val := string(template[i:i+2]) + "-40"
		fmt.Printf("%s %v\n", string(template[i:i+2]), count[val])
		for k, v := range count[val] {
			pairCount[k] += int(v)
			for _, c := range k {
				finalCount[string(c)] += int(v) 
			}
		}
		// 	finalCount = getMapSummation(finalCount, count[val], string(template[i:i+1]))
	}
	finalCount2 := make(map[string]int)
	// fmt.Printf("Pair count: %v\n", pairCount)
	for k, v := range pairCount {
		for _, c := range k {
			finalCount2[string(c)] += v / 2
		}
	}
	for k := range finalCount {
		if finalCount[k] % 2 == 1 {
			fmt.Printf("Not even %s\n", k)
		}
		finalCount[k] /= 2
	}
	fmt.Printf("Pair count: %v\n", pairCount)
	fmt.Printf("Final count: %v\n", finalCount)
	fmt.Printf("Final count2: %v\n", finalCount2)
	fmt.Printf("Answer: %d\n", getAnswer(finalCount))
}

func getCount(str string, step int, pairs map[string]string, countMap map[string]map[string]int64) map[string]int64 {
	if val, ok := countMap[str+"-"+fmt.Sprint(step)]; ok {
		return val
	}
	left := str[:1] + pairs[str]
	right := pairs[str] + str[1:]
	leftCount := getCount(left, step-1, pairs, countMap)
	rightCount := getCount(right, step-1, pairs, countMap)
	count := getMapSummation(leftCount, rightCount, pairs[str])
	countMap[str+"-"+fmt.Sprint(step)] = count
	// fmt.Printf("Str: %s, Left: %s, Right: %s, Step %d, Count: %v\n", str, string(left), string(right), step, count)
	return count
}

func getMapSummation(a, b map[string]int64, mid string) map[string]int64 {
	res := make(map[string]int64, 0)
	for k, v := range a {
		res[k] = v
	}
	for k, v := range b {
		res[k] += v
	}
	// res[mid] -= getMin(a[mid], b[mid])
	return res
}

func getMin(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func getAbs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func part1() {
	template := "NN"
	pairs := inputPairs

	byteTemp := []byte(template)
	steps := 40
	count := countInitial(template)
	fmt.Printf("Initial count: %v\n", count)
	for step := 0; step < steps; step++ {
		var buf bytes.Buffer
		for i := 0; i < len(byteTemp)-1; i++ {
			buf.WriteByte(byteTemp[i])
			searchStr := string(byteTemp[i : i+2])
			if val, ok := pairs[searchStr]; ok {
				buf.WriteString(val)
				count[val]++
			} else {
				panic("didn't find pair match")
			}
		}
		buf.WriteByte(byteTemp[len(byteTemp)-1])
		byteTemp = buf.Bytes()
		// fmt.Printf("Cur after step %d: %s\n", step+1, string(byteTemp))
		fmt.Printf("Step: %v\n", step+1)
		fmt.Printf("Count: %v\n", count)
	}
	// fmt.Printf("Answer: %d\n", getAnswer(count))
}

func getAnswer(input map[string]int) int {
	max := math.MinInt
	min := math.MaxInt
	for _, v := range input {
		max = findMax(max, v)
		min = findMin(min, v)
	}
	fmt.Printf("Max: %d, Min: %d\n", max, min)
	return max - min
}

func countInitial(str string) map[string]int64 {
	res := make(map[string]int64, 0)
	// for _, c := range str {
	// 	res[string(c)]++
	// }
	res[str] = 1
	return res
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var testTemplate = "NNCB"

var testPairs = map[string]string{
	"CH": "B",
	"HH": "N",
	"CB": "H",
	"NH": "C",
	"HB": "C",
	"HC": "B",
	"HN": "C",
	"NN": "C",
	"BH": "H",
	"NC": "B",
	"NB": "B",
	"BN": "B",
	"BB": "N",
	"BC": "B",
	"CC": "N",
	"CN": "C",
}

var inputTemplate = "VPPHOPVVSFSVFOCOSBKF"

var inputPairs = map[string]string{
	"CO": "B",
	"CV": "N",
	"HV": "H",
	"ON": "O",
	"FS": "F",
	"NS": "S",
	"VK": "C",
	"BV": "F",
	"SC": "N",
	"NV": "V",
	"NC": "F",
	"NH": "B",
	"BO": "K",
	"FC": "H",
	"NB": "H",
	"HO": "F",
	"SB": "N",
	"KP": "V",
	"OS": "C",
	"OB": "P",
	"SH": "N",
	"BC": "H",
	"CK": "H",
	"SO": "N",
	"SP": "P",
	"CF": "P",
	"KV": "F",
	"CS": "V",
	"FF": "P",
	"VS": "V",
	"CP": "S",
	"PH": "V",
	"OP": "K",
	"KH": "B",
	"FB": "S",
	"CN": "H",
	"KS": "P",
	"FN": "O",
	"PV": "O",
	"VC": "S",
	"HF": "N",
	"OC": "O",
	"PK": "V",
	"KC": "C",
	"HK": "C",
	"PO": "N",
	"OO": "S",
	"VH": "N",
	"CC": "K",
	"BP": "K",
	"HC": "K",
	"FV": "K",
	"KF": "V",
	"VF": "C",
	"HN": "S",
	"VP": "B",
	"HH": "O",
	"FO": "O",
	"PC": "N",
	"KK": "C",
	"PN": "P",
	"NN": "C",
	"FH": "N",
	"VV": "O",
	"OK": "V",
	"CB": "N",
	"SN": "H",
	"VO": "H",
	"BB": "C",
	"PB": "F",
	"NF": "P",
	"KO": "S",
	"PP": "K",
	"NO": "O",
	"SF": "N",
	"KN": "S",
	"PS": "O",
	"VN": "V",
	"SS": "N",
	"BF": "O",
	"HP": "H",
	"HS": "N",
	"BS": "S",
	"VB": "F",
	"PF": "K",
	"SV": "V",
	"BH": "P",
	"FP": "O",
	"CH": "P",
	"OH": "K",
	"OF": "F",
	"HB": "V",
	"FK": "V",
	"BN": "V",
	"SK": "F",
	"OV": "C",
	"NP": "S",
	"NK": "S",
	"BK": "C",
	"KB": "F",
}
