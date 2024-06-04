package main

import (
	"fmt"
	"slices"
)

type LeveledNode struct {
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

	leveledNodes := make([]LeveledNode, len(graph))
	levelGraphNodes(graph, 0, 1, []int{}, leveledNodes)

	fmt.Printf("%+v\n", leveledNodes)
	// Output
	// [{id:0 level:1} {id:1 level:2} {id:2 level:2} {id:3 level:3} {id:4 level:4}]

}

func levelGraphNodes(graph map[int][]int, current int, level int, visited []int, leveledNodes []LeveledNode) {
	graphEdges := graph[current]
	visited = append(visited, current)

	if leveledNodes[current].level < current {
		leveledNodes[current] = LeveledNode{id: current, level: level}
	}

	for _, edge := range graphEdges {
		if slices.Contains(visited, edge) {
			continue
		}

		levelGraphNodes(graph, edge, level+1, visited, leveledNodes)
	}
}
