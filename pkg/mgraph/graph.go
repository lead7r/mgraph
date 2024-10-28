package mgraph

import (
	"errors"
	"fmt"
	"maps"
)

var (
	ErrVertexAlreadyAdded  = errors.New("vertex already added")
	ErrVertexDoesNotExists = errors.New("vertex does not exists")
	ErrEdgeAlreadyAdded    = errors.New("edge already added")
)

type Graph interface {
	AddVertex(id VertexID) (Vertex, error)
	Vertex(id VertexID) Vertex
	RemoveVertex(id VertexID)
	Vertices() []Vertex
	ForEachVertex(func(v Vertex) bool)
	AddEdge(id EdgeID, from, to VertexID) (Edge, error)
	Edge(id EdgeID) Edge
	RemoveEdge(id EdgeID)
	Edges() []Edge
	ForEachEdge(func(e Edge) bool)
	EdgesBetween(from, to VertexID) []Edge
	Order() int
	Size() int
	Degree() int
	EnsureConsistency(enable bool)
	EnsuresConsistency() bool
	IsConsistent() bool
	Clone() Graph
}

func New() Graph {
	return &graph{
		properties: defaultProperties(),
		vertices:   make(map[VertexID]*vertex),
		edges:      make(map[EdgeID]*edge),
		edgesFrom:  make(map[VertexID]map[EdgeID]*edge),
		edgesTo:    make(map[VertexID]map[EdgeID]*edge),
	}
}

func defaultProperties() properties {
	return properties{
		consistency:              true,
		alwaysEnsuredConsistency: true,
	}
}

type graph struct {
	properties properties
	vertices   map[VertexID]*vertex
	edges      map[EdgeID]*edge
	edgesFrom  map[VertexID]map[EdgeID]*edge
	edgesTo    map[VertexID]map[EdgeID]*edge
}

type properties struct {
	consistency              bool
	alwaysEnsuredConsistency bool
}

func (g *graph) AddVertex(id VertexID) (Vertex, error) {
	if g.vertices[id] != nil {
		return nil, fmt.Errorf("error while adding vertex '%s': %w", id, ErrVertexAlreadyAdded)
	}
	v := newVertex(id, g)
	g.vertices[id] = v
	return v, nil
}

func (g *graph) Vertex(id VertexID) Vertex {
	if v, ok := g.vertices[id]; ok {
		return v
	}
	return nil
}

func (g *graph) RemoveVertex(id VertexID) {
	if g.properties.consistency {
		for id := range g.edgesFrom[id] {
			g.RemoveEdge(id)
		}
		for id := range g.edgesTo[id] {
			g.RemoveEdge(id)
		}
	}
	g.vertices[id].onRemove()
	delete(g.vertices, id)
}

func (g *graph) Vertices() []Vertex {
	vertices := make([]Vertex, 0, len(g.vertices))
	for _, v := range g.vertices {
		vertices = append(vertices, v)
	}
	return vertices
}

func (g *graph) ForEachVertex(each func(v Vertex) bool) {
	for _, v := range g.vertices {
		if !each(v) {
			return
		}
	}
}

func (g *graph) AddEdge(id EdgeID, from, to VertexID) (Edge, error) {
	if g.edges[id] != nil {
		return nil, fmt.Errorf("error while adding edge '%s': %w", id, ErrEdgeAlreadyAdded)
	}
	if g.properties.consistency {
		if g.vertices[from] == nil {
			return nil, fmt.Errorf("error while adding edge '%s': %w", id, fmt.Errorf("vertex '%s' does not exists: %w", from, ErrVertexDoesNotExists))
		}
		if g.vertices[to] == nil {
			return nil, fmt.Errorf("error while adding edge '%s': %w", id, fmt.Errorf("vertex '%s' does not exists: %w", to, ErrVertexDoesNotExists))
		}
	}
	e := newEdge(id, from, to, g)
	g.edges[id] = e
	if _, ok := g.edgesFrom[from]; !ok {
		g.edgesFrom[from] = make(map[EdgeID]*edge)
	}
	g.edgesFrom[from][id] = e
	if _, ok := g.edgesTo[to]; !ok {
		g.edgesTo[to] = make(map[EdgeID]*edge)
	}
	g.edgesTo[to][id] = e
	return e, nil
}

func (g *graph) Edge(id EdgeID) Edge {
	if e, ok := g.edges[id]; ok {
		return e
	}
	return nil
}

func (g *graph) RemoveEdge(id EdgeID) {
	edge := g.edges[id]
	delete(g.edgesFrom[edge.from], id)
	if len(g.edgesFrom[edge.from]) == 0 {
		delete(g.edgesFrom, edge.from)
	}
	delete(g.edgesTo[edge.to], id)
	if len(g.edgesTo[edge.to]) == 0 {
		delete(g.edgesTo, edge.to)
	}
	edge.onRemove()
	delete(g.edges, id)
}

func (g *graph) Edges() []Edge {
	edges := make([]Edge, 0, len(g.edges))
	for _, e := range g.edges {
		edges = append(edges, e)
	}
	return edges
}

func (g *graph) ForEachEdge(each func(e Edge) bool) {
	for _, e := range g.edges {
		if !each(e) {
			return
		}
	}
}

func (g *graph) EdgesBetween(from, to VertexID) (edges []Edge) {
	for _, e := range g.edgesFrom[from] {
		if e.to == to {
			edges = append(edges, e)
		}
	}
	return
}

func (g *graph) Order() int {
	return len(g.vertices)
}

func (g *graph) Size() int {
	return len(g.edges)
}

func (g *graph) Degree() (degree int) {
	for _, v := range g.vertices {
		if d := v.Degree(); d > degree {
			degree = d
		}
	}
	return
}

func (g *graph) EnsureConsistency(enable bool) {
	if !enable {
		g.properties.alwaysEnsuredConsistency = false
	}
	g.properties.consistency = enable
}

func (g *graph) EnsuresConsistency() bool {
	return g.properties.consistency
}

func (g *graph) IsConsistent() bool {
	if g.properties.alwaysEnsuredConsistency {
		return true
	}
	for _, e := range g.edges {
		if g.vertices[e.from] == nil || g.vertices[e.to] == nil {
			return false
		}
	}
	return true
}

func (g *graph) Clone() Graph {
	newG := &graph{
		properties: g.properties,
	}
	newG.vertices = make(map[VertexID]*vertex, len(g.vertices))
	newG.edges = make(map[EdgeID]*edge, len(g.edges))
	newG.edgesFrom = make(map[VertexID]map[EdgeID]*edge, len(g.edgesFrom))
	newG.edgesTo = make(map[VertexID]map[EdgeID]*edge, len(g.edgesTo))
	maps.Copy(newG.vertices, g.vertices)
	maps.Copy(newG.edges, g.edges)
	for id, from := range g.edgesFrom {
		newG.edgesFrom[id] = make(map[EdgeID]*edge, len(from))
		maps.Copy(newG.edgesFrom[id], from)
	}
	for id, to := range g.edgesTo {
		newG.edgesTo[id] = make(map[EdgeID]*edge, len(to))
		maps.Copy(newG.edgesTo[id], to)
	}
	return newG
}
