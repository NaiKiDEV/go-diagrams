package main

import (
	"fmt"
	"slices"
)

type LeveledNodes struct {
	id    int
	level int
}

func main() {
	graph := make(map[int][]int)

	graph[0] = []int{1, 2}
	graph[1] = []int{4}
	graph[2] = []int{3}
	graph[3] = []int{4}
	graph[4] = []int{}

	leveledNodes := make([]LeveledNodes, len(graph))
	levelGraphNodes(graph, 0, 1, []int{}, leveledNodes)

	fmt.Printf("%+v\n", leveledNodes)

}

func levelGraphNodes(graph map[int][]int, current int, level int, visited []int, leveledNodes []LeveledNodes) {
	graphEdges := graph[current]
	visited = append(visited, current)

	leveledNodes[current] = LeveledNodes{id: current, level: level}

	for _, edge := range graphEdges {
		if slices.Contains(visited, edge) {
			continue
		}

		levelGraphNodes(graph, edge, level+1, visited, leveledNodes)
	}
}
