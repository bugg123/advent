package main

import "fmt"

func main() {
	part1()
	part2()
}

func part1() {
	steps := 100
	input := inputs
	flashes := 0
	printBoard(input)
	for step := 1; step <= steps; step++ {
		flashed := make(map[int]map[int]bool)
		for row := 0; row < len(input); row++ {
			for col := 0; col < len(input[row]); col++ {
				if flashed[row][col] {
					continue
				}
				input[row][col]++
				if input[row][col] > 9 {
					flashes += flash(row, col, input, flashed)
				}
			}
		}
		fmt.Printf("Step: %d Flashes: %d\n", step, flashes)
	}
}

func part2() {
	input := inputs
	printBoard(input)
	step := 1
	for {
		flashed := make(map[int]map[int]bool)
		for row := 0; row < len(input); row++ {
			for col := 0; col < len(input[row]); col++ {
				if flashed[row][col] {
					continue
				}
				input[row][col]++
				if input[row][col] > 9 {
					flash(row, col, input, flashed)
				}
			}
		}
		if allFlashed(input, flashed) {
			fmt.Printf("All flashes: %d\n", step)
			break
		}
		step++
	}
}

func allFlashed(input [][]int, flashed map[int]map[int]bool) bool {
	if len(input) != len(flashed) {
		return false
	}
	for i := 0; i < len(input); i++ {
		if len(input[i]) != len(flashed[i]) {
			return false
		}
	}
	return true
}

func printBoard(input [][]int) {
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			print(input[row][col])
		}
		println()
	}
	println()
}

func flash(row, col int, input [][]int, flashed map[int]map[int]bool) int {
	if flashed[row][col] || row < 0 || col < 0 || row >= len(input) || col >= len(input[row]) {
		return 0
	}
	input[row][col]++
	if input[row][col] <= 9 {
		return 0
	}

	if _, ok := flashed[row]; !ok {
		flashed[row] = make(map[int]bool)
	}
	flashed[row][col] = true
	input[row][col] = 0
	res := 1
	res += flash(row-1, col-1, input, flashed)
	res += flash(row-1, col, input, flashed)
	res += flash(row-1, col+1, input, flashed)
	res += flash(row, col+1, input, flashed)
	res += flash(row, col-1, input, flashed)
	res += flash(row+1, col+1, input, flashed)
	res += flash(row+1, col, input, flashed)
	res += flash(row+1, col-1, input, flashed)
	return res
}

var testInputs = [][]int{
	{1, 1, 1, 1, 1},
	{1, 9, 9, 9, 1},
	{1, 9, 1, 9, 1},
	{1, 9, 9, 9, 1},
	{1, 1, 1, 1, 1},
}

var testInputs2 = [][]int{
	{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
	{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
	{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
	{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
	{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
	{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
	{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
	{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
	{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
	{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
}
var inputs = [][]int{
	{8, 2, 7, 1, 6, 5, 3, 8, 3, 6},
	{7, 5, 6, 7, 6, 2, 6, 7, 7, 5},
	{2, 3, 1, 5, 7, 1, 3, 3, 1, 6},
	{6, 5, 4, 2, 6, 5, 5, 3, 1, 5},
	{2, 4, 5, 3, 6, 3, 7, 3, 3, 3},
	{1, 2, 4, 7, 2, 6, 4, 3, 2, 8},
	{2, 3, 2, 5, 1, 4, 6, 6, 1, 4},
	{2, 1, 1, 5, 8, 4, 3, 1, 7, 1},
	{6, 1, 8, 2, 3, 7, 6, 2, 8, 2},
	{2, 3, 8, 4, 7, 3, 8, 6, 7, 5},
}
