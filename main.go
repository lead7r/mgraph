package main

import (
	"errors"
	"fmt"
	"graph/pkg/mgraph"
)

func main() {
	g := mgraph.New()
	v, err := g.AddVertex("node1")
	if err != nil {
		panic(err)
	}
	v.StoreData(user{})
	v2, err := g.AddVertex("node2")
	if err != nil {
		panic(err)
	}
	e, err := g.AddEdge("edge1", "node1", "node2")
	if err != nil {
		panic(err)
	}
	fmt.Printf("vertices obtenidos: %v %v\n", v, v2)
	fmt.Printf("edge obtenido: %v\n", e)
	g.ForEachVertex(func(v mgraph.Vertex) bool {
		fmt.Printf("vertice recorrido: %v\n", v)
		return true
	})

	edges := g.EdgesBetween("node1", "node2")
	if len(edges) > 0 {
		edges[0] = nil
	}
	fmt.Printf("obtenidos %v\n", edges)
	fmt.Printf("nuevamente %v\n", g.EdgesBetween("node1", "node2"))

	_, err = g.AddEdge("edge1", "node1", "node2")
	fmt.Println(err)
	if errors.Is(err, mgraph.ErrEdgeAlreadyAdded) {
		fmt.Printf("el error es un vertex added\n")
	}
}
