package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := inputs
	incomplete := part1(input)
	score := map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}
	scores := make([]int, 0)
	for _, str := range incomplete {
		stack := GetIncStack(str)
		totalScore := 0
		l := len(stack)
		for i := 0; i < l; i++ {
			val := stack.Pop()
			totalScore = totalScore*5 + score[val]
		}
		scores = append(scores, totalScore)
	}
	sort.Ints(scores)
	fmt.Printf("Part 2 Middle Score: %d\n", scores[len(scores)/2])
}

func part1(input []string) []string {
	score := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	notCorrupted := make([]string, 0)
	corrupted := make([]string, 0)
	totalPoints := 0
	for _, str := range input {
		c := CheckCorruptedString(str)
		if points, ok := score[c]; ok {
			corrupted = append(corrupted, str)
			totalPoints += points
		} else {
			notCorrupted = append(notCorrupted, str)
		}
	}
	fmt.Printf("Part 1 Total Score: %d\n", totalPoints)
	return notCorrupted
}

func GetIncStack(str string) Stack {
	rightToLeft := map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
	var stack Stack
	for _, c := range str {
		if _, ok := rightToLeft[c]; ok {
			if err := stack.Check(rightToLeft[c]); err != nil {
				panic("Corrupted in incomplete strings")
			}
		} else {
			stack.Push(c)
		}
	}
	return stack
}

func CheckCorruptedString(str string) rune {
	rightToLeft := map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
	var stack Stack
	for _, c := range str {
		if _, ok := rightToLeft[c]; ok {
			if err := stack.Check(rightToLeft[c]); err != nil {
				return c
			}
		} else {
			stack.Push(c)
		}
	}
	return 0
}

type Stack []rune

func (s *Stack) String() string {
	str := ""
	for _, v := range *s {
		str += string(v) + " "
	}
	return strings.TrimSpace(str)
}

func (s *Stack) Check(toCheck rune) error {
	if len(*s) == 0 {
		return errors.New("Check on empty stack")
	}
	val := s.Pop()
	if val != toCheck {
		return errors.New("Invalid match")
	}
	return nil
}

func (s *Stack) Push(v rune) {
	*s = append(*s, v)
}

func (s *Stack) Pop() rune {
	if len(*s) == 0 {
		return 0
	}
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

var testInputs = []string{
	"[({(<(())[]>[[{[]{<()<>>",
	"[(()[<>])]({[<{<<[]>>(",
	"{([(<{}[<>[]}>{[]{[(<()>",
	"(((({<>}<{<{<>}{[]{[]{}",
	"[[<[([]))<([[{}[[()]]]",
	"[{[{({}]{}}([{[{{{}}([]",
	"{<[[]]>}<{[{[{[]{()[[[]",
	"[<(<(<(<{}))><([]([]()",
	"<{([([[(<>()){}]>(<<{{",
	"<{([{{}}[<[[[<>{}]]]>[]]",
}

var inputs = []string{
	"({<{({([<<{<[<<><>><<>>](<()()>)><{{{}<>}}<<<>{}>[<>[]]>>}{[<<[]{}>[<>()]>[(()[])]]<{{<><>}[[]<>]}{<",
	"[(<[[([[(([<{{[]{}}([][])}<<{}[]>([]{})>><{<()()><[][]>}{[()]<()[]>}>]<{({()()}){[{}[]][<>()]}}[{[{}<>]",
	"{{(<[[(<<<(<<[{}{}]{<>()}><{[]{}}([]{})>>([([][]){{}<>}][<()<>>[(){}]]))<{<(<>())[{}<>]>{{[]{}}<<>{}>}}({<",
	"<([({[<[(([[<(()<>){[]<>}><<(){}>{()()}>]](<(({}[]){[][]})<{<>())[<>[]]>>{{{()<>}{(){}}}{[<>()]<{}[]>}",
	"<(((<{<(<{{[[[()<>]{[]()}]][{<()[]><<>{}>}]}}[<{[{[][]}[<><>]]}{{{(){}}[<>()]}{[[]()]{<>[]}}}>{[({{}",
	"{{({[({(({{[<[{}()]([]<>)><<{}()>[<>[]]>]<<{[]{}}([]<>)>{{<>}{<>{}}}}}{{<<[]<>>{<>()}>}<({[]<>}",
	"(([{[(<(<[[(([()[]][{}()])(([][])<{}<>>))]]>)[<(<{{<<><>>(<>[])}<([]<>)[[]{}]>}((<{}[]>[<><>]){{[",
	"{{[<<{<<(<([<<[]()>[[][]]>{[()[]]<()[]>}]){(<(<>{})([])>{<(){}>{<>()]}){[<[]()><<>()>](([][",
	"((<{[{{(<<[(<{<>{}}[<>[]]>{<{}<>>{{}[]}})]>><{([<(<>())[{}()]>{<<>{}>{[]{}}}><([[]()]<[]<>>)(<[]>{()()})>)}>)",
	"{<({{{(<[([<{(<>[])<()[]>}([[]]([]{}))>])[{{[([]){[]{}}]}}{[[{(){}}{<>{}}][[(){}]]]<([<>{}]{{}<>})(",
	"([([{{{[[<<[[(()[])<<><>>]{({}[])}]((([]{})<<>>)({[]<>}[<>[]]))>{(({<>()})([{}[]]<()<>>)){[",
	"{[<(({[<<{{{[({}[]){<>()}]{[()()]{[]()}}}[{<<><>>{<>[]}}<{()()}<()>>]}({<(<><>)[()[]]>([{}()])}{[(<>())(()",
	"<{{[[<{{(<<{<[{}{}]<{}[]>>[{{}[]}[{}()]]}((<()()>{(){}})<((){})>)>[{<({})[[]{}]>([{}<>])}{{[{}()]}{<[",
	"[{[<{<[[({([(<{}<>>[<>()])]<<<{}{}>({}<>)>>)<{(({}[])({}<>))}<{<[]>({}<>)}{{()()}([]())}>>}){<[<[[{}<>]<",
	"{(<{{<{<<(<<<[{}{}]((){})>>{<<{}{}>{()[]}>}>)([({(<>())<<>()>>)(<<()[]>(<>())><<()()>({}())>)]{(<[<><",
	"[{<{[<{([{[({<(){}>([]())}[<<>[]>([]())]){{<(){}>{{}[]}}<[()[]]({})]}](((<{}()>{{}()})))}[[(",
	"<<[{{[<<(<{({{[]{}}{()()}}){<<{}()>({}<>)>(({}<>)[()<>])}}<[[([]{})[<><>]]]((<()()>{<>[]}){{[]{",
	"[[({[({[([({{<[]{}>{(){}}}{[[][]][()()]}}[{<{}><{}()>}({()[]})]){{[[()<>]<[][]>][({}<>)]}}](<({<{}(",
	"[[[((([((<{{[(<>())[{}[]]]{{<>{}}}}}>))<[[[({{{}{}}(()())})({([]{})[<><>]}[([]{}){()[]}])]<{<[{",
	"[{[<[{{([[{(({()()}<{}[]>){{()<>}{{}()}})<<<{}[]>>{<<>[]><{}{}>}>}]{{{[{()()}[<>()}]{{{}}[(){}]}}}",
	"[[{[(({[[[[{(<<>()>(<>()))[<{}[]>{[]<>}]}{{{()<>}<()()>}}]<[<([]<>}{<>[]}>]>](((<{()()}<()()>><<<>[]><<>",
	"({[{({(<{((<<[()[]]<()<>>]<([][])>>[([{}()]<{}{}>){<{}{}>(()[])}])[{([<>()]{()[]}){[()[]]}}([{()<>}",
	"[(({(([[{(<{{<[][]>(()()}}[({}[])]}[(<<>[]><()[]>)]>)<{({{[]{}}[()[]]}[<<>()>(()())])}{(((()[])<[]>)",
	"<(<([[({[((((<{}[]>[{}{}]){[{}<>](<><>)})(([()<>]({})){[{}{}]{[]<>}}))[({(()[])<<>[]>}[[()<>]<()",
	"{[(<<({([{[(<<[][]>(<>())>){<[<>()][[]{}]>(({}[])[(){}])}][(({{}<>}))(<[()()]<()()>>}]}{[{<(<>",
	"[{{({<({<(<[[<{}<>>{()[]}][{()()}]]>){[<{{()[]}[{}()]}>([<()[]><{}<>>](<()[]>))]{[<[<><>]{<>[]}>[<()",
	"{[<<[{<<[[(<(<[][]>[()()])>){<(([]()){{}[]})><<[{}()]<<>>>>}}[[([<{}[]><(){}>][({}())[()<>]])[[{<>[]",
	"{((<[{{<{[<[((()<>)<()()>)]>]}>}((([{[[<<>[]>[[][]]]<[[]()]({}[])]]{[([]{}){()()}]}}<((<[]<",
	"{[[({<{<<{<{{{[]{}}}}>}[[{(([][]){<>[]}){{()<>}}}<<<()[]>[{}[]]>>]<[{[()[]]({}{})}{[{}()]}]>]>>}>}<",
	"[{[(<({{(<(([{()<>}{()<>}]{{[]}(()())]))>(<<{[<><>]<()<>>}[(<>{})<<>{}>]>[{{<>{}}{{}[]}}<[<",
	"{{((<[(([<{{([()][(){}])({<>()}[()<>])}[{[{}<>]<()()>}<{{}}([]<>)>>}(<<[[]{}]<[]<>>>[([]()){[]{}}]>((<<>(",
	"(<<<(<([<<({<<<><>>[{}<>]>{<()>([]<>)}}{{(<><>)}<{<>()}{{}()}>})<<<({}{})<[]<>>>>>]{({({[][]}[{}()])([<>()](",
	"([(<<<(<(([([{<>{}}<(){}>](<{}<>>[{}{}]})][([<()>{{}[]}](<<>()><()[]>))[[{<>()}[<>[]]]{<<>[]>",
	"(<{[[<<<({[<<([]())[()()]>([()[]]{<>{}})>]})[[{<[{<>()}({}{})><(()<>)([]{})>><<<<>()><<>()",
	"[[[{(<[[((<[{<<>[]>({}[])}<[{}{}]{[]<>}>]<{({}{})}{[{}<>]{<><>}}>>([{[[]<>](()<>)}{{<>{}}(<>[]))]{[{{}{}}]})",
	"[(<([[[{<[[<((()()){[][]})>((<{}[]>[()<>])<([]()){<>[]}>)]<<{<<><>>}{{(){}}}>>]<{[<<(){}>({}<>)><<<>()>([]<>",
	"[<((<<[<{[{[{[()[]]{()<>}}(<<>()>[<>()])]}{<[[(){}]{(){}}]{[()]{<>[]}}>{{<[][]>[()()]}[(<>{",
	"[({{([([{[[([<<>{}>[[]()]]({{}()}))][<{{()[]}}<<{}{}>>><[[{}[]][<>()]][{{}[]}(<>{}))>]][{((((){",
	"((<<{(([[[{{<[()[]]({}())>}<[{(){}}]<({}{})<()[]>>>}{((<()<>>([]<>)){[()][()[]]})([{()}<[]{}>]<(<><>",
	"<<<<[([<<(({<[[]()]<<><>>>({[][]})}{([<>{}]<<>()>){{[]{}}}})){[[[[[]{}][(){})]{(()[])(()<>)}]{{[[]<>][{}(",
	"[({{([<{({{[{(()[])<<><>>}(((){})[<>()])>}{{{{{}<>}<{}[]>}{[(){}](()[])}}}}([(([<>{}](<>[]))",
	"([(<({{(<<[([[<>][<>()]])<[{()[]}([]{})]{[()<>][[]{}]}>]><<{[{<><>}<()()>]{{<>[]}>}>>><[<{{",
	"({{(([({{<[<[{{}{}}{{}<>}]((<>{})[[]<>])>(((<><>)<{}()>){{{}{}}{{}{}}})]({<{{}}([][])><<[]{}>",
	"[(<{({{(<<<(<[{}[]]{(){}}>(([]()){<>[])))[{{[]{}}{()()}}{[<>{}][<>]}]>[((<{}<>>[<>[]])[({}[])<[]{}>])[<[{",
	"[(<<<[<<({{{[[(){}]]{<<>{}>{()[]}}}{({<>{}}{[][]})}}{<([()[]]<{}<>>)([{}()]({}[]))>[(<<>()>{[]()})]}}{([{{(",
	"{[{{(({[{<{{[[()<>]<[]<>>]({{}{}}[{}{}])}{(<<>[]>{[]})[[[]{}]{()<>}>}}{<<{{}{}}(()[])>>{{[[]()]",
	"[<[[{(([<<[<[([][])<[]{}}]{<(){}>}>(<<()<>><[][]>><<{}[]>[{}()]>)]><([([[]<>](()[]))[(()<>)<(){}>]][({",
	"<{{<([<([{{{(<[]{}>[[]{}])[{<>{}}([])]}}[<[<<><>>({}<>)]{((){})[[]<>]}><[{()[]}<<>{}>]{{()()",
	"{[<<<([{([{<<(()[]){{}{}}>[(<>{}){{}()}]><[[{}<>][[]()]](<{}()>([]()))>}]){{<{<[()<>][<>()]><",
	"<{[<<[([[{[({[[]<>](<>[])})({([]<>)<[]()>}<([]<>)[<>[]]>)]}[[{[{()<>}{<>{}}][<[]{}>(<>[])]}<({[]<>",
	"[[[[([({(((<<[()<>]>[(()())]>{[{{}()}[[]()]]{[<>()]([][])}})))}(<({(({()[]}[<>()]))<[[()<>]]<{{}()}[()<>]",
	"{[([([<[(<<[<[[]()]<[]<>>>][{<[][]>[<>()]}{({}<>)<()[]>}]><{({<>()}{<><>})<{(){}}({}{})>}<{(",
	"<<({[<<{(<[<({{}[]}{<>[]})[<()()>]>(([<>()][<><>])(<<>{}><()[]>))](<[<{}<>>[{}[]]]({{}{}}(",
	"[<<{<(([<[{(([()<>]({}{})){<<>[]>[()()]})({([])(<>())}[{[][]}[{}[]]])}{({(()())[()[]]}({[]<>}([][])))",
	"((({<[{([(<[(<[]()>)](<[{}()]<[]{}>>(({}())<()[]>))>)[(<[([]<>)(<>{})]>)]][({({[()[]]{<>}}[[[]{}]{()[]}])([<<",
	"[[[(({[<[((({([]){<>{}}})[[{()[]}[{}[]]][{[]<>}]])<<(<[]{}>[<>[]])[<[][]>[{}[]]]>>)[([{{<><>}[[][]]",
	"({<(({{[{{<[((<>())(<>[]))<<<><>>((){})>]([<[][]><{}{}>][[{}()]<<>()>])>((<[[]()]([][])>[<{}[]>{",
	"(({<<([[<[[{<[()()]<[][]}>[{(){}}[{}<>]]}(<{{}<>}{[]<>}>{<(){}>{<><>}})]<<{<<>[]>[<>[]]}{{",
	"<[{([<[({<((<[{}()]([]<>))[(()[])])((({}{})<()<>>){(()<>)}))>{[<<[<>{}][(){}]>{<<><>>(<>[])}>{[",
	"{[([[[{{<<[<<<()()>(()<>)>[({}{})[<>{}]]>([[[]<>]]{<<>>[[][]]})]><<[{(()[])][<()>[{}[]]]]<([{}{}][()<>])(",
	"{(({{<<((<{<<<[]{}>{{}()}><<<>()><()[]>>><<[<>[]]](({}[])[[]{}])>}[[({<>()}{()}){[[]{}]({}())}",
	"<[[{<[({<([<<<()<>>>>([[[]()]{<><>}](<<>[]>[<>]))])(({({[]{}}){{{}()}{(){}}}}[{(()<>)[[]()]}[{<>()}{[][]}]]))",
	"{<[[([{<<<{{[<[]{}>[[]<>]]{({}())<[][]>}}<<<{}{}>>({(){}})>}([{({}())}<([]())[()()]>])>({{[[<>()]{{}{",
	"<{(<[{[[{[<<<{(){}}[<>{}]>{{{}[]}<<>()>}>{[{[]()}<<>>]{<[]{}>[()()]}}>][[[{<<>())[[]<>]}<(<>",
	"[{{<<{{([(<<{{{}}[()]}<(()())(<>())>>(<{[]()}[[]<>]>)>[((<()()><[]{}>}({[]<>}))(({[][]}{[]<>})<(()[])(()[]",
	"([(<<[<{({[[[(()<>)<{}()>]{[{}()]{<>{}}}]{(<(){}>{{}[]}){<<><>><{}[]>}}]}<{[[[()()]{{}{}}]]",
	"{<(((<({<((([<(){}>]){{{<><>}{{}()}}})({<[[]{}]>{({}())}}))>})([[<<<{[{}[]]{[]{}}}<{()<>}({}[])>>>{[",
	"{([[{[[[<<<({{()<>}}(<[]<>>[<><>]))[[[[]<>][{}()]]]>{<<[<>()][[][]]](<{}{}>{[]{}})>[<(<>{}){<><>}>([{}[]]",
	"[[{(<{{{<(<{[[{}()](()[])]}<<{(){}}[()[]]>{[()()]([]())}>>)>[(<<<{[][]}>>[{([][])[(){}]}[(<",
	"<{[(<{<{(<<{<{()[]}(()[])>(([]<>)([]()))}>{((({}[]))<{[]()}{{}()>>)<{<(){}>}<{<>()}<()()>>",
	"<(<[{{[[{<{[{[<>][<>[]]}]{[(<>[])][<{}{}>{<>}]}}(({[()<>]{[]<>}}<{(){}}[[]{}]>)([[<><>]<()()>]<({}[])([]())",
	"{(<{[{[{({{(<<<>{}>><({}())(<>())>)}}{(<<<(){}>{(){}}>(<{}()>(<>()))><[{<>()}{{}()}][{[]<>}(()())]>){<[{{}[]",
	"{[((((({<(([(<[]<>><[]<>>)<<{}()>>](<{()[]}(()<>)>{{{}{}}([]{})})))[<[{([])[<>]}(([]())<{}{}>]]{",
	"(<{[((([{(<<{(()<>){{}[]}}>><<<({}<>)[[]()]>{(<>[])[()()]}>{{{<><>}([]())}}>)[{(([[]()][<>(",
	"<[<<<[(<[[(<<([]{})>([[][]][{}{}])><[<<>[]>[{}[]]]<[<>][()()]>>)][[<<{()[]}(<><>)><<()()>[()<>]>>][{[{[]()",
	"<<{([[<[[<[({[<>[]](()<>)}{({}[]){()[]}}){({{}[]}(<>()))}]<[[<[][]>[[]<>]]<[{}<>]<[]{}>>)>>]<",
	"[((([[{<<[[<{{{}[]}[()<>]}>(<<{}{}>[(){}]>({<>[]}<{}<>>))][<{[{}<>]{[]}}>({[{}{}}[[][]]})]](({[{[][]}]{[{}",
	"[([{<[[{<((<<{[]<>}(()<>)>>{(<<>{}>(()<>))}))[(<[[[][]][<>{}]]{[()[]](()())}><({{}{}}}[({}())]>",
	"{{[({<(({{[({{[]()}<()<>>}{([]()){[]{}}}){[<<>>[<><>]][<{}<>><()[]>]}][{<({}{}><[][]>>{{{}()}<(){}>}}]}}))>}<",
	"{[{[<{{({[({[[{}()][(){}]]{{[][]}<<>()>}}[[(<>[])<[]{}>]((<>[])[{}()])])]})}}>{({{<([{<([]())[{}{}]>{",
	"<{{[{<{<<{{({<{}[]>[(){}]}[[<>()]{(){}}])}}[{{[(()<>)(<><>)]<{()<>}{<>()}>}{({[]()}{{}()})<{{",
	"((((([{{<<[<({()<>}<<>{}))<{{}<>}(<>[])>>{<{(){}}[{}[]]>}]{<<[{}{}]>{{()()}[{}<>]}>}>>([[({(()())<()",
	"(({{[[<[[<<<<<[]()>[{}<>]]>>({[<<><>><{}()>]{{(){}}[<>{}]}}[[([]<>)({}<>)]])>{<[[<{}[]><{}<>>][<<>[]>]](<<[]",
	"<[<([<<[([[[[<{}[]><()>][(<>{}){[]}]]{({()[]}[<>{}])}]([((<>[])(<>[]))[{<>[]}{[]<>}]][(([]{})({}<>",
	"<((<{{({{([<([<>{}](()()))<<(){}>[<>[]>>>({<<>>([]{})})]<[[([]())<<>>]{{<>{}}(()<>)}]>)<[(((<>)<{}<>>)({<>(",
	"{({<({<<[[{[((<>())<()<>>)<(<>[])((){})>]({{[]()}[{}()]}[[{}{}]<{}[])])}({<[(){}](<><>)><{()<>}[{}[]",
	"<[(<<[((([(([[[]<>][<><>]]{{{}{}}[{}[]]}){(({}())<[]{}>)})<(<(()[]){()()}>{<()>[[]{}]}){{[[]{}]([]<",
	"{(<(([{[[<[[[{<>{}}{[]{}}]({(){}})]<[[()()]][[{}[]]]>]{<<{{}{}}[{}{}]>>[(<{}()>([]{}))(<(){}>{<><>}",
	"((<<({({<{<<[<[]{}><()[]>]<(<>()){()[]}>>>}{(((({}<>)({}()))<{[]{}}<<>[]>>)<({[][]}<(){}>){",
	"[[<{((({{{(<{[[]{}}{{}}}<{[]{}}<()>>>{<{{}[]}<[]()>>{[()[]]{[][]}}})([{[[][]]}(([]{})<<><>>)][(<[][]>{{}",
	"(<({[{[(<<(((([]())[<><>]}([<><>]))<[[()<>]<<><>>]{<[]()>{()<>}}>)[<<{[]()}(())>><[<(){}><<>{}",
	"(<<[<(<{<({<[{(){}}{<><>}][<<>()>]>}<[<[()[]]<{}[]>>]<<({}[]]><[[][]]<[]()>>>>)>}><[((<[<<[]<>>>[[<>()][{}(",
	"[((<[{[[([{[[<<>>{<>()}]<[[]{}]<[]<>>>]}<({(()<>))<<<>()><<>[]>>)<{<[]{}>((){})}[({}[]){()()}]>",
	"<<{{({{([{{<<{{}{}}(<>{})><{()()}[{}<>]>><{<[]<>>[[]()]}[<<><>>({})]>}(((<<>{}><{}{}>)([<>{}][{}]))[{",
}
