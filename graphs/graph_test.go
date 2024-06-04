package graphs

import (
	"testing"
)

type BasicNode struct {
	title string
}

func NewBasicNode(title string) BasicNode {
	return BasicNode{title: title}
}

func TestGraph(t *testing.T) {
	graph := New[BasicNode]()

	vertex1 := NewNode(1, NewBasicNode("User"))
	vertex2 := NewNode(2, NewBasicNode("Contacts"))
	vertex3 := NewNode(3, NewBasicNode("Address"))

	graph.AddNode(vertex1)
	graph.AddNode(vertex2)
	graph.AddNode(vertex3)

	graph.AddEdge(vertex1, vertex2)
	graph.AddEdge(vertex1, vertex3)
	graph.AddEdge(vertex2, vertex3)

	vertex4 := NewNode(4, NewBasicNode("Testing"))
	graph.AddNode(vertex4)
	graph.AddEdge(vertex4, vertex2)
	graph.AddEdge(vertex4, vertex3)
	graph.AddEdge(vertex4, vertex1)
	graph.AddEdge(vertex4, vertex1)

	vertex5 := NewNode(5, NewBasicNode("SAME"))
	graph.AddNode(vertex5)
	graph.AddEdge(vertex5, vertex5)
	graph.AddEdge(vertex5, vertex5)
	graph.AddEdge(vertex5, vertex5)

	edges := graph.GetEdges()

	for _, vertex := range graph.GetNode() {
		vertexEdges := make([]string, 0)
		for _, edge := range edges {
			if vertex == edge.GetFromVertex() {
				vertexEdges = append(vertexEdges, edge.GetToVertex().GetData().title)
				// Debug info
				t.Logf("'%s' -> '%s'\n", vertex.GetData().title, edge.GetToVertex().GetData().title)
			}
		}

		// Need recursive traversal for deeper levels
		if len(vertexEdges) == 0 {
			continue
		}
	}
}
