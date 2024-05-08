package main

import (
	"fmt"
	"strings"

	"github.com/NaiKiDEV/go-diagrams/graphs"
)

type BasicNode struct {
	title string
}

func NewBasicNode(title string) BasicNode {
	return BasicNode{ title: title }
}

func main() {
	graph := graphs.New[BasicNode]();
	
	vertex1 := graphs.NewVertex(1, NewBasicNode("User"))
	vertex2 := graphs.NewVertex(2, NewBasicNode("Contacts"))
	vertex3 := graphs.NewVertex(3, NewBasicNode("Address"))

	graph.AddVertex(vertex1)
	graph.AddVertex(vertex2)
	graph.AddVertex(vertex3)
	
	graph.AddEdge(vertex1, vertex2)
	graph.AddEdge(vertex1, vertex3)
	graph.AddEdge(vertex2, vertex3)

	vertex4 := graphs.NewVertex(4, NewBasicNode("Testing"))
	graph.AddVertex(vertex4)
	graph.AddEdge(vertex4, vertex2)
	graph.AddEdge(vertex4, vertex3)
	graph.AddEdge(vertex4, vertex1)
	graph.AddEdge(vertex4, vertex1)
	
	vertex5 := graphs.NewVertex(5, NewBasicNode("SAME"))
	graph.AddVertex(vertex5)
	graph.AddEdge(vertex5, vertex5)
	graph.AddEdge(vertex5, vertex5)
	graph.AddEdge(vertex5, vertex5)
		
	edges := graph.GetEdges()
	for _, vertex := range graph.GetVertexes() {
		vertexEdges := make([]string, 0)
		for _, edge := range edges {
			if vertex == edge.GetFromVertex() {
				vertexEdges = append(vertexEdges, edge.GetToVertex().GetData().title)
				// Debug info
				// fmt.Printf("'%s' -> '%s'\n", vertex.GetData().title, edge.GetToVertex().GetData().title)
			}
		}

		// Need recursive traversal for deeper levels
		
		if len(vertexEdges) == 0 {
			continue
		}

		// Rendering logic for each vertex (1 level support only)
		vertexEdgeTitleLength := 0
		for _, edgeTitle := range vertexEdges {
			vertexEdgeTitleLength += len(edgeTitle)
		}
		
		vertexTitle := vertex.GetData().title
		vertexTitleLength := len(vertexTitle)
		
		if len(vertexEdges) > 1 {
			defaultGap := 1
			edgeLengthWithDefaultGap := vertexEdgeTitleLength + defaultGap * (len(vertexEdges))
			vertexTitleEdgeTitleDiff := edgeLengthWithDefaultGap - vertexTitleLength
			vertexTitlePaddingLeft := (edgeLengthWithDefaultGap - vertexTitleLength) / 2

			gapModifier := 0
			if vertexTitleEdgeTitleDiff % 2 == 1 {
				gapModifier = 1
			}

			calculatedGap := defaultGap + gapModifier

			fmt.Println(strings.Repeat(" ", vertexTitlePaddingLeft), vertexTitle)
			for _, edgeTitle := range vertexEdges {
				fmt.Print(edgeTitle, strings.Repeat(" ", calculatedGap))
			}
			fmt.Println()
		} else if len(vertexEdges) == 1 {
			fmt.Println(vertexTitle)
			fmt.Println(strings.Repeat(" ", vertexTitleLength / 2 - vertexEdgeTitleLength % 2) + "|")
			fmt.Println(vertexEdges[0])
		}

		fmt.Println();
	}
}