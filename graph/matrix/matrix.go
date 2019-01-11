package matrix

import (
	"fmt"

	"github.com/altay13/gostructures/hashtable"

	"github.com/altay13/gostructures/set"

	"github.com/altay13/gostructures/queue"
)

// Graph is the implementation of Graph
// Adjacency Matrix
type Graph struct {
	matrix     [][]int
	size       int
	queue      *queue.Queue
	visitedSet *set.Set
}

// NewGraph ...
func NewGraph(size int) *Graph {
	mtrx := make([][]int, size)

	for i := 0; i < size; i++ {
		mtrx[i] = make([]int, size)
	}

	s := &Graph{
		matrix:     mtrx,
		size:       size,
		queue:      queue.NewQueue(),
		visitedSet: set.NewSet(size),
	}

	return s
}

// AddVertice ...
func (g *Graph) AddVertice() {
	g.size++
	g.matrix = append(g.matrix, make([]int, g.size))
	for i := 0; i < g.size; i++ {
		g.matrix[i] = append(g.matrix[i], 0)
	}
}

// AddEdge ...
func (g *Graph) AddEdge(v1, v2 int) {
	g.matrix[v1][v2] = 1
	g.matrix[v2][v1] = 1
}

// Print method prints graph matrix
func (g *Graph) Print() {
	fmt.Printf("  ")
	for i := 0; i < g.size; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")
	for i := 0; i < g.size; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < g.size; j++ {
			fmt.Printf("%d ", g.matrix[i][j])
		}
		fmt.Printf("\n")
	}
}

// GetNeighbours method returns all neighbours of v
func (g *Graph) GetNeighbours(v int) []int {
	res := make([]int, 0)
	for i := 0; i < g.size; i++ {
		if g.matrix[v][i] == 1 {
			res = append(res, i)
		}
	}

	return res
}

// DFS method performs Depth First Search
func (g *Graph) DFS(v int, route *[]int) {
	*route = append(*route, v)
	g.visitedSet.Add(v)
	for i := 0; i < g.size; i++ {
		if g.matrix[v][i] == 1 && !g.visitedSet.Contains(i) {
			g.DFS(i, route)
		}
	}
}

// BFS method performs Breadth First Search
func (g *Graph) BFS(v int, route *[]int) {
	if g.visitedSet.Contains(v) {
		return
	}
	*route = append(*route, v)
	g.visitedSet.Add(v)
	for i := 0; i < g.size; i++ {
		if g.matrix[v][i] == 1 && !g.visitedSet.Contains(i) {
			g.queue.Enqueue(i)
		}
	}

	for !g.queue.IsEmpty() {
		g.BFS(g.queue.Dequeue().(int), route)
	}
}

// ShortestPath method search for shortest path
// with BFS
func (g *Graph) ShortestPath(v, t int, route *[]int) {
	g.queue.RemoveAll()
	dist := hashtable.NewHash(g.size)
	pred := hashtable.NewHash(g.size)

	for i := 0; i < g.size; i++ {
		dist.Add(i, -1)
	}

	res := g.bfsShortestPath(v, t, dist, pred)

	if res {
		c := t
		*route = append(*route, c)
		for pred.Get(c) != nil && pred.Get(c).(int) != -1 {
			*route = append(*route, pred.Get(c).(int))
			c = pred.Get(c).(int)
		}
	}
}

func (g *Graph) bfsShortestPath(v, t int, dist, pred *hashtable.Hash) bool {
	if v == t {
		return true
	}
	if g.visitedSet.Contains(v) {
		return false
	}

	g.visitedSet.Add(v)
	for i := 0; i < g.size; i++ {
		if g.matrix[v][i] == 1 && !g.visitedSet.Contains(i) {
			g.queue.Enqueue(i)
			dist.Add(i, dist.Get(v).(int)+1)
			pred.Add(i, v)
		}
	}

	for !g.queue.IsEmpty() {
		return g.bfsShortestPath(g.queue.Dequeue().(int), t, dist, pred)
	}

	return false
}

// CountIslands ...
func (g *Graph) CountIslands() int {
	count := 0

	for row := 0; row < g.size; row++ {
		for col := 0; col < g.size; col++ {
			if g.matrix[row][col] == 1 && !g.visitedSet.Contains(col) {
				count++
			}
		}
	}

	return count
}

// DisposeVisitedSet removes all edges from visitedSet
func (g *Graph) DisposeVisitedSet() {
	g.visitedSet.RemoveAll()
}
