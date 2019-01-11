package main

import (
	"fmt"

	"github.com/altay13/gostructures/graph/matrix"
)

func main() {
	g := matrix.NewGraph(8)

	g.AddEdge(0, 1)
	g.AddEdge(0, 3)
	g.AddEdge(1, 2)
	g.AddEdge(3, 4)
	g.AddEdge(3, 7)
	g.AddEdge(4, 5)
	g.AddEdge(4, 6)
	g.AddEdge(4, 7)
	g.AddEdge(5, 6)
	g.AddEdge(6, 7)
	g.Print()

	// lens := make([]int, 0)

	ttmp := make([]int, 0)
	g.BFS(2, &ttmp)
	fmt.Println(ttmp)

	g.DisposeVisitedSet()

	ttmp = make([]int, 0)
	g.ShortestPath(0, 7, &ttmp)

	fmt.Println(ttmp)
}
