package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Node struct {
	Name        string
	Connections map[string]*Node
}

type NodeConnections map[string]Node
type Path []*Node

func main() {
	connections := parseRoutes(inputs)
	start := connections["start"]
	testPaths := findPaths(&start, Path{}, connections, make(map[string]bool, 0), false)
	fmt.Printf("Found %v paths\n", len(testPaths))
}

func findPaths(n *Node, curPath Path, c NodeConnections, visited map[string]bool, secondUsed bool) []Path {
	if visit := visited[n.Name]; (visit && secondUsed) || (len(curPath) != 0 && n.Name == "start") {
			return nil 
	}
	if n.Name == "end" {
		curPath = append(curPath, n)
		return []Path{curPath}
		// curPath = curPath[:len(curPath)-1]
	}
	if unicode.IsLower([]rune(n.Name)[0]) {
		if visited[n.Name] {
			secondUsed = true
			defer func() {
				secondUsed = false
			}()
		} else {
			visited[n.Name] = true
			defer delete(visited, n.Name)
		}
	}

	curPath = append(curPath, n)
	defer func() {
		curPath = curPath[:len(curPath)-1]
	}()

	paths := make([]Path, 0)
	for _, node := range n.Connections {
		paths = append(paths, findPaths(node, curPath, c, visited, secondUsed)...)
	}
	return paths
}

var validPaths = make([]Path, 0)

func findPathsGlobal(n *Node, curPath Path, c NodeConnections, visited map[string]bool, secondUsed bool) {
	if visit := visited[n.Name]; (visit && secondUsed) || (len(curPath) != 0 && n.Name == "start") {
			return 
	}
	if n.Name == "end" {
		curPath = append(curPath, n)
		validPaths = append(validPaths, curPath)
		curPath = curPath[:len(curPath)-1]
		return
	}
	if unicode.IsLower([]rune(n.Name)[0]) {
		if visited[n.Name] {
			secondUsed = true
			defer func() {
				secondUsed = false
			}()
		} else {
			visited[n.Name] = true
			defer delete(visited, n.Name)
		}
	}

	curPath = append(curPath, n)
	defer func() {
		curPath = curPath[:len(curPath)-1]
	}()

	for _, node := range n.Connections {
		findPaths(node, curPath, c, visited, secondUsed)
	}
}

func findPaths1(n *Node, curPath Path, c NodeConnections, visited map[string]bool) []Path {
	if visit, ok := visited[n.Name]; ok || visit {
		return nil
	}
	if n.Name == "end" {
		curPath = append(curPath, n)
		return []Path{curPath}
	}
	if unicode.IsLower([]rune(n.Name)[0]) {
		visited[n.Name] = true
	}
	curPath = append(curPath, n)
	paths := make([]Path, 0)
	for _, node := range n.Connections {
		if _, ok := visited[node.Name]; ok {
			continue
		}
		paths = append(paths,findPaths1(node, curPath, c, visited)...)
	}
	curPath = curPath[:len(curPath)-1]
	delete(visited, n.Name)
	return paths
}

func parseRoutes(routes []string) NodeConnections {
	results := make(NodeConnections)
	for _, route := range routes {
		parts := strings.Split(route, "-")
		results.addConnections(parts[0], parts[1])
	}
	return results
}

func (c NodeConnections) addConnections(from, to string) {
	if _, ok := c[from]; !ok {
		c[from] = Node{from, make(map[string]*Node)}
	}
	if _, ok := c[to]; !ok {
		c[to] = Node{to, make(map[string]*Node)}
	}
	toNode := c[to]
	toAddr := &toNode
	c[from].Connections[to] = toAddr
	fromNode := c[from]
	fromAddr := &fromNode
	c[to].Connections[from] = fromAddr
}

func (c NodeConnections) String() string {
	var result string
	for _, node := range c {
		result += node.String() + "\n"
	}

	return result
}

func (n Node) String() string {
	return fmt.Sprintf("%s", n.Name)
}

var inputs = []string{
	"end-ry",
	"jf-jb",
	"jf-IO",
	"jb-hz",
	"jo-LM",
	"hw-end",
	"hw-LM",
	"hz-ry",
	"WI-start",
	"LM-start",
	"kd-jf",
	"xi-WI",
	"hw-jb",
	"hz-jf",
	"LM-jb",
	"jb-xi",
	"ry-jf",
	"WI-jb",
	"end-hz",
	"jo-start",
	"WI-jo",
	"xi-ry",
	"xi-LM",
	"xi-hw",
	"jo-xi",
	"WI-jf",
}

var testInputs = []string{
	"start-A",
	"start-b",
	"A-c",
	"A-b",
	"b-d",
	"A-end",
	"b-end",
}

var testInputs2 = []string{
	"dc-end",
	"HN-start",
	"start-kj",
	"dc-start",
	"dc-HN",
	"LN-dc",
	"HN-end",
	"kj-sa",
	"kj-HN",
	"kj-dc",
}

var testInputs3 = []string{
	"fs-end",
	"he-DX",
	"fs-he",
	"start-DX",
	"pj-DX",
	"end-zg",
	"zg-sl",
	"zg-pj",
	"pj-he",
	"RW-he",
	"fs-DX",
	"pj-RW",
	"zg-RW",
	"start-pj",
	"he-WI",
	"zg-he",
	"pj-fs",
	"start-RW",
}
