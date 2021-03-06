package main

import "fmt"

func main() {
	fmt.Printf("Ans: %d\n", getTreeCount(3, 1))
	ans := getTreeCount(1, 1) * getTreeCount(3, 1) * getTreeCount(5, 1) * getTreeCount(7, 1) * getTreeCount(1, 2)
	fmt.Printf("Final Ans: %d\n", ans)
}

func getTreeCount(right, down int) int {
	count := 0
	col := 0
	for i := 0; i < len(inputs); i += down {
		row := inputs[i]
		if row[col] == '#' {
			count++
		}
		if col+right >= len(row) {
			col = col - len(row)
		}
		col += right
	}

	return count
}

var testInputs = []string{
	"..##.........##.........##.........##.........##.........##.......",
	"#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..",
	".#....#..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.",
	"..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#",
	".#...##..#..#...##..#..#...##..#..#...##..#..#...##..#..#...##..#.",
	"..#.##.......#.##.......#.##.......#.##.......#.##.......#.##.....",
	".#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#",
	".#........#.#........#.#........#.#........#.#........#.#........#",
	"#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...",
	"#...##....##...##....##...##....##...##....##...##....##...##....#",
	".#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#",
}

var inputs = []string{
	".....##.#.....#........#....##.",
	"....#...#...#.#.......#........",
	".....##.#......#.......#.......",
	"...##.........#...#............",
	"........#...#.......#.........#",
	"..........#......#..#....#....#",
	"..................#..#..#....##",
	".....##...#..#..#..#..#.##.....",
	"..##.###....#.#.........#......",
	"#.......#......#......#....##..",
	".....#..#.#.......#......#.....",
	"............#............#.....",
	"...#.#........#........#.#.##.#",
	".#..#...#.....#....##..........",
	"##..........#...#...#..........",
	"...........#...###...#.......##",
	".#..#............#........#....",
	"##.#..#.....#.......#.#.#......",
	".##.....#....#.#.......#.##....",
	"..##...........#.......#..##.#.",
	"##...#.#........#..#...#...#..#",
	".#..#........#.#.......#..#...#",
	".##.##.##...#.#............##..",
	"..#.#..###......#..#......#....",
	".#..#..#.##.#.##.#.#...........",
	"...#....#..#.#.#.........#..#..",
	"......#.#....##.##......#......",
	"#....#.##.##....#..#...........",
	"...#.#.#.#..#.#..#.#..#.##.....",
	"#.....#######.###.##.#.#.#.....",
	"..#.##.....##......#...#.......",
	"..#....#..#...##.#..#..#..#..#.",
	".............#.##....#.........",
	".#....#.##.....#...............",
	".#............#....#...#.##....",
	".#.....#.##.###.......#..#.....",
	".#...#.........#.......#..#....",
	"..#.#..#.##.......##...........",
	".....##..#..#..#..#.##..#.....#",
	"..##............##...#..#......",
	"...#..#....#..##.....##..#.#...",
	"#.....##....#.#.#...#...#..##.#",
	"#.#..#.........#.##.#...#.#.#..",
	".....#.#....##....#............",
	"#.......#..#.....##..#...#...#.",
	".....#.#...#...#..#......#.....",
	"..##....#.#.#.#.#..#...........",
	"##..#...#.........#......#...#.",
	"..#...#.#.#.#..#.#.##..##......",
	"#............###.....###.......",
	"..........#...#........###.....",
	".......##...#...#...#........#.",
	".#..#.##.#.....................",
	".#..##........##.##...#.......#",
	".......##......#.....#......#..",
	".##.#.....#......#......#......",
	"#...##.#.#...#.#...............",
	"........#..#...#.##.......#....",
	"...................#...#...##..",
	"...#...#.........#.....#..#.#..",
	".###..#........#..##.##..#.##..",
	"#...#.....#.....#.....#..#..#..",
	"###..#.....#.#.#.#......#....#.",
	"#........#....##.#...##........",
	".#.#..##........##....##.#.#...",
	"#...#....#.###.#.#.........#...",
	"...#...##..###.......#.........",
	"......#....#..##..#.....#.#....",
	"........#...##...###......##...",
	"..........##.#.......##........",
	"...#....#......#...##.....#....",
	"###.#.....#.#..#..#....#...#..#",
	".#.....#.#....#...............#",
	"..#....#....####....###....#.#.",
	"....##........#..#.##.#....#...",
	".......##...#...#..#....####...",
	"#...##.#......##...#..#........",
	"..##..#.##....#.......##.#.#...",
	"..#.#...............#...#.#....",
	"....#.....#.#.....#.##.......#.",
	"...#.#..##.#.#..............##.",
	"..#.....#...#.............#.##.",
	"##..#.#...#........#..#.....##.",
	"...........##...#.#.###...#....",
	"...#.#.#..#..................#.",
	".#...##.............#...#......",
	"..#..#...#.#.......#...#.....#.",
	"..##.......#.#.................",
	".##..#........###.....#....#.##",
	"......#..###.......#....##....#",
	"....#.....#.................#..",
	"........#...#...#..............",
	"...#..#.###.......#..#.#.#.##..",
	"..#...#.....#....#.........#...",
	"...#.............#........###..",
	"......#..............#......#..",
	"#..#...........#...#..........#",
	"...##...#.###..#...#.....#.#...",
	"....#..##......#.......##......",
	"....#....##.#...#.#..#....#...#",
	".#...........#..#....##...#..##",
	"..#.#.................###.#...#",
	"..#.#.#...##...........#.......",
	"..........#..##...#.#..##....##",
	"........#........#.##..#.#...#.",
	".....#...##.......##......#...#",
	"....#...#..#..#.....#..........",
	".#..#......#..#..#..###.......#",
	".##..........#...#...#.#.....##",
	"..#..........#.#.#...###.......",
	"....#................#...##....",
	".##..#....#..........#.#.#.....",
	"..##...#.#........#.....#.##...",
	"....####.....#..#.........##..#",
	"......#.........#...#..........",
	"....#...................#..##..",
	".##....#.#.........#....#...#..",
	"....##...##.....#..####........",
	"..##.#....#.#.......##...#.....",
	"#...#.#.#...#..#..##.....#.....",
	"#..................###.....#...",
	"#.#.....#.......#.#...###.#....",
	".#..#....#............#........",
	"#.#....#..#.#...............#..",
	"..#..#..#.............#......#.",
	"..#.......##...................",
	".#....#.........#....#.#.#..#..",
	"....#....#..#...............#..",
	"......#..#..##......#.........#",
	"..#.##........##......#..#..#.#",
	"#.....#.#....#.........##...#..",
	"###..............#....###...##.",
	"....#..##......#.......##......",
	"......#...#.##......##....#..#.",
	"..........#....#..##.......#..#",
	".#..#...##..#...........#..#..#",
	".....#....#...#..###...###....#",
	".#####..#...#.#.#..#.#.###...##",
	"..##............##.#...#.##...#",
	".##..#...#...#....##.#..#..##..",
	".#....#...#............##..#...",
	".#.#......#....#....#..##..##..",
	".........#...#.......#.##..#...",
	"#.........#.....##.....#..#..#.",
	"...##.#...#...#..#..#....##..##",
	".#............#...#....##......",
	"..#...#.##.........#.#......#.#",
	"....#.##........#.........#..##",
	"#.........#......#.#......#..#.",
	"........#.#.......#.#........#.",
	"..#..........##.#...#..#.#.....",
	"..#...#....#...#...#..#.#..#.#.",
	".#.........#....#..#####..#....",
	"#.#....#.#.###...#.............",
	"..##...........##......##......",
	"#.....#..#....#...............#",
	"...#.#..#....##......#...##....",
	"...#........#.....#...#..#.....",
	".#......##.........#......#....",
	"..#..###.##...#.#.....#........",
	".............#......#..#.......",
	"..#...............#.#...#..#..#",
	".......#..#...#.#####......#..#",
	".........#.....#...............",
	"##........#............#.#.....",
	".#...#.....#..#..#...#....#...#",
	"..#....#....##......##.....#.#.",
	"#...##..##......#...#....#.....",
	"....#.#.#.....###....##.##....#",
	"..........##...##.......#......",
	"..#.......#...##.#....##.##....",
	"....#........................#.",
	"...#...#.#.##...#.....#...#..#.",
	".#....##..#..#..........##..##.",
	".#.....#..#...#.##.....#.......",
	".#.##...#.#..#.....##....#...#.",
	".##...#........##....#..#......",
	".....#........#..........#.#..#",
	"....#..##.......#..#.....#.....",
	"...........#...#........#.##..#",
	".....#..#....#..#.#.....#....##",
	".....#....#.##.#..##...........",
	"...##.......##.........#.......",
	"...............##..#....#.#....",
	".......###..#........#..####.##",
	".......#.##...#.#....#.####....",
	"....#...............#..........",
	"##.#.......#.....#......#...#..",
	"......##.....#....#.....#..#..#",
	".....#...##.............#......",
	"#.#.##.#.....#..#........#.....",
	"......##....#..#........#......",
	"............#........#..#.#....",
	"##.......#......#...####..#.##.",
	"..##..#...#.............#.##...",
	".....#..##......#.##......###..",
	"............#........#........#",
	"#.#.#.#...#.#.....#.........#..",
	".........#...............#.....",
	".............###.#.......#....#",
	"###.##..#..#..........#....#...",
	"#......#...#..#..#.....#.##....",
	"............#....#....#..#.....",
	"..#.#....#...#......#.#..#..##.",
	"...#........................#..",
	"#.#...#..........#......#.#....",
	".........#................#...#",
	"##.....#....#........##.......#",
	"#...##........#...#...........#",
	"...#...#..........##.......#.#.",
	"..#.#.#....#......##...........",
	"...#.#...#.##.#..#.#.##........",
	"#....##.....###..#.......#.....",
	"###.....#.#.#...#..#.........##",
	"..#......#..###...#.#.#.....#.#",
	".#....#.....#............#..##.",
	"....#....##..........#.....##..",
	"#...........#....#...#..#...##.",
	"..#.......#.....#..........#...",
	".#..#................#......#..",
	"..#......#.#...#..#.#....#....#",
	"...#..#...###..#..##....#.#....",
	"..#..............#.....#.......",
	"...#.#...#.........#.#.........",
	"##......##...........##.#.##..#",
	"..#..##..#....#.#......#.#...##",
	"...#.###....###...#.....#......",
	"#.#................#......#....",
	"..#.....#.....#....##.......#..",
	".#.#...............##..#.......",
	"...#....#.......#.#.....##..#..",
	".........#....#.......#.#...##.",
	"#....#......##.#.........##...#",
	"#.............#..##.#.#..##....",
	"...#....#..#...#....#.#.#.#...#",
	".#....#....#..##.....#.#...###.",
	"##............#.#...##.#..#.#..",
	"##.#....##.....#..#..###....#..",
	"##....#................##......",
	"...##..#...#..###....#.....##..",
	".#...##......#..#.#.....#...#..",
	"..##......##...#.##.......#....",
	"......#.....#.....##........#.#",
	"##....#...........#............",
	"#.......#....#..#.##..##.#..#..",
	".#....##.#.....#..#..#.........",
	".#....#.#.#...#.....##.....#.#.",
	".......##.#.#........#......##.",
	"##........#.##.......#...#..#..",
	"...###..##....#.#....#.#.......",
	"......#.......#...##.....#...#.",
	"..#......##.#......#.....#.....",
	".....#.....###...#.............",
	"#...#.#...#...#..#......#......",
	"#.....#.......###.#....###.#...",
	"...#.......#....####....##..#..",
	"#.#.....#....#........#.......#",
	".........#.......#......#.#...#",
	"..##....#.....##...............",
	"..........#..#.#..#......#.....",
	"..................##...##.#....",
	"........#.......#...#..#.#.#...",
	".....#.#..##..#..#.#..#.......#",
	".....#........#..#..#....#....#",
	"##............#..#..#...#....#.",
	".....#....................##..#",
	"........##.#....###............",
	"##.......#.##................#.",
	".....###.#..#..#...#....###.##.",
	".#......#.#....#.....##.#......",
	"...##......##.........#...#....",
	"....####..............#........",
	"#...#.#..##..##.........##.....",
	"......#......#....#..#.........",
	"#.....#.....#.##...............",
	"..#.##..#...##.#.####..#....###",
	"#..#......#....#.##..##...#.#..",
	"#....#.......#.....#.....#.#...",
	"##.......#.....##...#.....#....",
	"...#...##..........#..##..##..#",
	".###..#..##...#....#...#..#....",
	"......##..###.......###...#....",
	"....#...#.#.......#.##...##..##",
	"#.#......#..##.#.#..#..#..#....",
	"......#........#.......#.......",
	"..........#.#.....##...........",
	"......#..#........#..#.#..###..",
	"##..#.............##..#........",
	".........#....#.....#.........#",
	".....#..##...#..#..##.##......#",
	"###..#...........#.......#....#",
	"...............#....#.#........",
	".##.#...#.#........##....#.....",
	".##.###...##..###....#...#...#.",
	".##..#....#.#.#...#.#.#.#...#..",
	".###.#...#.......#....#..#.....",
	"..#..#.#.#.#........#.....##...",
	".#.......#.#...#.#...........##",
	"...#.....##....#.....##...#....",
	"................#.....####...#.",
	".#.#......#.......##...#.##....",
	".###.........#.#......#..#.#...",
	"#......#...#....#..##.......#..",
	".##..#....#..#...........#...#.",
	".#...#.......##........#.##....",
	"..#...........#...##...........",
	".....##....##......#....#..#...",
	"#......#.#...#.##.#...##....#..",
	"#....................#...##...#",
	"..#............#........#......",
	".............#.........##.....#",
	"...#...#......##.#...#...#.#...",
	"..#...#.#.................#....",
	"....##...#....#...###.##......#",
	"...#....#...#..#...#....#.....#",
	"...##.#........#..#.........#..",
	"..##.....#..##...#.....##...#..",
	"#.........#.#.#...#......#...#.",
	"#.#...........#...#..#..#..##..",
	"..#..#..##....#..........#.###.",
	".....#..#....#.#...#...#..#..#.",
	"###.....#..#.................#.",
	".#..##.##.#......#....##..#....",
}
