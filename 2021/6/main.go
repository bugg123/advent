package main

import "fmt"

var memo map[string]int

func main() {
	memo = make(map[string]int)
	counter := len(input)
	for _, i := range input {
		counter += getFish(i, 0, 256)
	}
	fmt.Println(counter)
}

func getFish(timer, curDay, lastDay int) int {
	if val, ok := memo[fmt.Sprintf("%d-%d", timer, curDay)]; ok {
		return val
	}
	if curDay >= lastDay {
		return 0
	}
	if timer == 0 {
		val := 1 + getFish(6, curDay+1, lastDay) + getFish(8, curDay+1, lastDay)
		memo[fmt.Sprintf("%d-%d", timer, curDay)] = val
		return val
	} else {
		val := getFish(timer-1, curDay+1, lastDay)
		memo[fmt.Sprintf("%d-%d", timer, curDay)] = val
		return val
	}
}

var input = []int{4, 5, 3, 2, 3, 3, 2, 4, 2, 1, 2, 4, 5, 2, 2, 2, 4, 1, 1, 1, 5, 1, 1, 2, 5, 2, 1, 1, 4, 4, 5, 5, 1, 2, 1, 1, 5, 3, 5, 2, 4, 3, 2, 4, 5, 3, 2, 1, 4, 1, 3, 1, 2, 4, 1, 1, 4, 1, 4, 2, 5, 1, 4, 3, 5, 2, 4, 5, 4, 2, 2, 5, 1, 1, 2, 4, 1, 4, 4, 1, 1, 3, 1, 2, 3, 2, 5, 5, 1, 1, 5, 2, 4, 2, 2, 4, 1, 1, 1, 4, 2, 2, 3, 1, 2, 4, 5, 4, 5, 4, 2, 3, 1, 4, 1, 3, 1, 2, 3, 3, 2, 4, 3, 3, 3, 1, 4, 2, 3, 4, 2, 1, 5, 4, 2, 4, 4, 3, 2, 1, 5, 3, 1, 4, 1, 1, 5, 4, 2, 4, 2, 2, 4, 4, 4, 1, 4, 2, 4, 1, 1, 3, 5, 1, 5, 5, 1, 3, 2, 2, 3, 5, 3, 1, 1, 4, 4, 1, 3, 3, 3, 5, 1, 1, 2, 5, 5, 5, 2, 4, 1, 5, 1, 2, 1, 1, 1, 4, 3, 1, 5, 2, 3, 1, 3, 1, 4, 1, 3, 5, 4, 5, 1, 3, 4, 2, 1, 5, 1, 3, 4, 5, 5, 2, 1, 2, 1, 1, 1, 4, 3, 1, 4, 2, 3, 1, 3, 5, 1, 4, 5, 3, 1, 3, 3, 2, 2, 1, 5, 5, 4, 3, 2, 1, 5, 1, 3, 1, 3, 5, 1, 1, 2, 1, 1, 1, 5, 2, 1, 1, 3, 2, 1, 5, 5, 5, 1, 1, 5, 1, 4, 1, 5, 4, 2, 4, 5, 2, 4, 3, 2, 5, 4, 1, 1, 2, 4, 3, 2, 1}

var testInput = []int{3, 4, 3, 1, 2}
