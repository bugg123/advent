package main

import "testing"

func BenchmarkFindPaths(b *testing.B) {
	connections := parseRoutes(inputs)
	start := connections["start"]
	for i := 0; i < b.N; i++ {
		findPaths(&start, Path{}, connections, make(map[string]bool, 0), false)
	}

}
func BenchmarkFindPathsGlobal(b *testing.B) {
	connections := parseRoutes(testInputs)
	start := connections["start"]
	for i := 0; i < b.N; i++ {
		findPathsGlobal(&start, Path{}, connections, make(map[string]bool, 0), false)
	}
}

