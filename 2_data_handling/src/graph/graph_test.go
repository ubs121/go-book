package graph

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGraphCreate(t *testing.T) {
	g := sampleGraph()
	exp := 3
	got := g.Degree(0)
	if got != exp {
		t.Errorf("exp %d, got %d", exp, got)
	}
}

func TestBFS(t *testing.T) {
	g := sampleGraph()

	// new bfs
	bfs := NewBFS(g, 2)
	if !bfs.HasPathTo(3) {
		t.Errorf("exp true, got false")
	}

	path := bfs.PathTo(3)
	fmt.Printf("2.pathTo(3) = %v", path)
}

func sampleGraph() *GraphAdjSet {
	g := NewGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(1, 3)
	return g
}

// a graph representation using a map
type Graph2 struct {
	data map[string]map[string]int
}

func NewGraph2() *Graph2 {
	g := &Graph2{}
	g.data = make(map[string]map[string]int)
	return g
}

// оройн жагсаалт
func (g *Graph2) V() []string {
	v := make([]string, 0)
	for k := range g.data {
		v = append(v, k)
	}
	return v
}

// ирмэг нэмэх
func (g *Graph2) addEdge(source, target string, weight int) {
	map2, ok := g.data[source]
	if !ok {
		map2 = make(map[string]int)
	}
	map2[target] = weight
	g.data[source] = map2
}

func (g *Graph2) String() string {
	var strBuf bytes.Buffer
	for k, map2 := range g.data {
		strBuf.WriteString(k + "\n")
		for k2, w := range map2 {
			strBuf.WriteString(fmt.Sprintf("  -> %s -> %d \n", k2, w))
		}
	}
	return strBuf.String()
}

func TestGraph(t *testing.T) {
	g := NewGraph2()
	g.addEdge("A", "B", 2)
	g.addEdge("A", "C", 0)
	g.addEdge("A", "G", 3)
	g.addEdge("A", "H", 5)
	g.addEdge("B", "A", 2)
	g.addEdge("C", "A", 0)
	g.addEdge("G", "A", 3)
	g.addEdge("G", "C", 0)
	g.addEdge("H", "A", 5)

	fmt.Println(g.V())
}
