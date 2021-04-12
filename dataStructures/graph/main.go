package main

import "fmt"

// we are implementing a directed graph
type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	Key      int
	Adjacent []*Vertex
}

func (g Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.Key)
		for _, vv := range v.Adjacent {
			fmt.Printf("%v", vv.Key)
		}
	}
}

func contains(s []*Vertex, val int) bool {
	for _, v := range s {
		if val == v.Key {
			return true
		}
	}
	return false
}

func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		e := fmt.Errorf("Invalid edge %v --> %v", from, to)
		fmt.Println(e.Error())
	} else if contains(fromVertex.Adjacent, to) {
		e := fmt.Errorf("Edge exists %v --> %v", from, to)
		fmt.Println(e.Error())
	} else {
		// add edge
		fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)
	}
}
func (g *Graph) getVertex(key int) *Vertex {
	for i, v := range g.vertices {
		if key == v.Key {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		e := fmt.Errorf("Error, vertex %v already exists", k)
		fmt.Println(e.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{Key: k})
	}
}
func main() {
	text := &Graph{}

	for i := 0; i < 5; i++ {
		text.AddVertex(i)
	}
	text.AddEdge(1, 2)
	text.AddEdge(1, 2)
	text.AddEdge(6, 2)
	text.AddEdge(3, 2)

	text.Print()
	// fmt.Println()

}
