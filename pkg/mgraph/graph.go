package mgraph

import (
	"errors"
	"fmt"
)

var (
	ErrVertexAlreadyAdded = errors.New("vertex already added")
	ErrEdgeAlreadyAdded   = errors.New("edge already added")
)

// el grafo no asegura consistencia
type Graph[V any, E any] interface {
	AddVertex(id VertexID) (Vertex[V], error)
	Vertex(id VertexID) Vertex[V]
	RemoveVertex(id VertexID)
	Vertices() []Vertex[V]
	ForEachVertex(func(v Vertex[V]) bool)
	AddEdge(id EdgeID, from, to VertexID) (Edge[E], error)
	Edge(id EdgeID) Edge[E]
	RemoveEdge(id EdgeID)
	Edges() []Edge[E]
	ForEachEdge(func(e Edge[E]) bool)
	EdgesBetween(from, to VertexID) []Edge[E]
	Incoming(id VertexID) []Edge[E]
	Outgoing(id VertexID) []Edge[E]
}

func New[V any, E any]() Graph[V, E] {
	return &graph[V, E]{
		vertices:  make(map[VertexID]*vertex[V]),
		edges:     make(map[EdgeID]*edge[E]),
		edgesFrom: make(map[VertexID]map[EdgeID]*edge[E]),
		edgesTo:   make(map[VertexID]map[EdgeID]*edge[E]),
	}
}

type graph[V any, E any] struct {
	vertices  map[VertexID]*vertex[V]
	edges     map[EdgeID]*edge[E]
	edgesFrom map[VertexID]map[EdgeID]*edge[E]
	edgesTo   map[VertexID]map[EdgeID]*edge[E]
}

func (g *graph[V, E]) AddVertex(id VertexID) (Vertex[V], error) {
	if g.vertices[id] != nil {
		return nil, fmt.Errorf("error while adding vertex '%s': %w", id, ErrVertexAlreadyAdded)
	}
	v := newVertex[V](id)
	g.vertices[id] = v
	return v, nil
}

func (g *graph[V, E]) Vertex(id VertexID) Vertex[V] {
	if v, ok := g.vertices[id]; ok {
		return v
	}
	return nil
}

func (g *graph[V, E]) RemoveVertex(id VertexID) {
	delete(g.vertices, id)
}

func (g *graph[V, E]) Vertices() []Vertex[V] {
	vertices := make([]Vertex[V], 0, len(g.vertices))
	for _, v := range g.vertices {
		vertices = append(vertices, v)
	}
	return vertices
}

func (g *graph[V, E]) ForEachVertex(each func(v Vertex[V]) bool) {
	for _, v := range g.vertices {
		if !each(v) {
			return
		}
	}
}

func (g *graph[V, E]) AddEdge(id EdgeID, from, to VertexID) (Edge[E], error) {
	if g.edges[id] != nil {
		return nil, fmt.Errorf("error while adding edge '%s': %w", id, ErrEdgeAlreadyAdded)
	}
	e := newEdge[E](id, from, to)
	g.edges[id] = e
	if _, ok := g.edgesFrom[from]; !ok {
		g.edgesFrom[from] = make(map[EdgeID]*edge[E])
	}
	g.edgesFrom[from][id] = e
	if _, ok := g.edgesTo[to]; !ok {
		g.edgesTo[to] = make(map[EdgeID]*edge[E])
	}
	g.edgesTo[to][id] = e
	return e, nil
}

func (g *graph[V, E]) Edge(id EdgeID) Edge[E] {
	if e, ok := g.edges[id]; ok {
		return e
	}
	return nil
}

func (g *graph[V, E]) RemoveEdge(id EdgeID) {
	edge := g.edges[id]
	delete(g.edgesFrom[edge.from], id)
	if len(g.edgesFrom[edge.from]) == 0 {
		delete(g.edgesFrom, edge.from)
	}
	delete(g.edgesTo[edge.to], id)
	if len(g.edgesTo[edge.to]) == 0 {
		delete(g.edgesTo, edge.to)
	}
	delete(g.edges, id)
}

func (g *graph[V, E]) Edges() []Edge[E] {
	edges := make([]Edge[E], 0, len(g.edges))
	for _, e := range g.edges {
		edges = append(edges, e)
	}
	return edges
}

func (g *graph[V, E]) ForEachEdge(each func(e Edge[E]) bool) {
	for _, e := range g.edges {
		if !each(e) {
			return
		}
	}
}

func (g *graph[V, E]) EdgesBetween(from, to VertexID) (edges []Edge[E]) {
	for _, e := range g.edgesFrom[from] {
		if e.to == to {
			edges = append(edges, e)
		}
	}
	return
}

func (g *graph[V, E]) Incoming(id VertexID) (edges []Edge[E]) {
	for _, e := range g.edgesTo[id] {
		edges = append(edges, e)
	}
	return
}

func (g *graph[V, E]) Outgoing(id VertexID) (edges []Edge[E]) {
	for _, e := range g.edgesFrom[id] {
		edges = append(edges, e)
	}
	return
}
