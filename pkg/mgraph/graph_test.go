package mgraph

import (
	"errors"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	g := New()
	if g == nil {
		t.Fatalf("New() returned nil")
	}
}

func TestGraph_AddVertex(t *testing.T) {
	g := New()
	v, err := g.AddVertex("node1")
	if err != nil {
		t.Fatalf("AddVertex() expected no error, got: %v", err)
	}
	if v == nil {
		t.Fatalf("AddVertex() returned nil vertex")
	}
}

func TestGraph_AddVertex_ErrorVertexAlreadyAdded(t *testing.T) {
	g := New()
	_, _ = g.AddVertex("node1")
	_, err := g.AddVertex("node1")
	if !errors.Is(err, ErrVertexAlreadyAdded) {
		t.Fatalf("AddVertex() expected error ErrVertexAlreadyAdded, got: %v", err)
	}
}

func TestGraph_Vertex(t *testing.T) {
	g := New()
	v := g.Vertex("node1")
	if v != nil {
		t.Fatalf("Vertex() expected to return nil vertex, got: %v", v)
	}
	_, _ = g.AddVertex("node1")
	v = g.Vertex("node1")
	if v == nil {
		t.Fatalf("Vertex() expected to return a vertex, got: %v", v)
	}
}

func TestGraph_RemoveVertex(t *testing.T) {
	g := New()
	v, _ := g.AddVertex("node1")
	g.RemoveVertex("node1")
	v2 := g.Vertex("node1")
	if v2 != nil {
		t.Fatalf("RemoveVertex() expected to remove vertex, got: %v", v2)
	}
	if v.Graph() != nil {
		t.Fatalf("RemoveVertex() expected to remove vertex from graph, got: %v", v.Graph())
	}
}

func TestGraph_RemoveVertexOnConsistentGraph_ShouldRemoveIncidentEdges(t *testing.T) {
	g := New()
	g.EnsureConsistency(true)
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	_, _ = g.AddVertex("node3")
	_, _ = g.AddVertex("node4")
	_, _ = g.AddEdge("edge1", "node1", "node2")
	_, _ = g.AddEdge("edge2", "node3", "node1")
	_, _ = g.AddEdge("edge3", "node3", "node4")
	g.RemoveVertex("node1")
	v := g.Vertex("node1")
	if v != nil {
		t.Fatalf("RemoveVertex() expected to remove vertex, got: %v", v)
	}
	e := g.Edge("edge1")
	if e != nil {
		t.Fatalf("RemoveVertex() expected to remove edge, got: %v", e)
	}
	e = g.Edge("edge2")
	if e != nil {
		t.Fatalf("RemoveVertex() expected to remove edge, got: %v", e)
	}
	e = g.Edge("edge3")
	if e == nil {
		t.Fatalf("RemoveVertex() expected to keep edge, got: %v", e)
	}
}

func TestGraph_RemoveVertexOnNonConsistentGraph_ShouldNotRemoveIncidentEdges(t *testing.T) {
	g := New()
	g.EnsureConsistency(false)
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	_, _ = g.AddVertex("node3")
	_, _ = g.AddVertex("node4")
	_, _ = g.AddEdge("edge1", "node1", "node2")
	_, _ = g.AddEdge("edge2", "node3", "node1")
	_, _ = g.AddEdge("edge3", "node3", "node4")
	g.RemoveVertex("node1")
	v := g.Vertex("node1")
	if v != nil {
		t.Fatalf("RemoveVertex() expected to remove vertex, got: %v", v)
	}
	e := g.Edge("edge1")
	if e == nil {
		t.Fatalf("RemoveVertex() expected to keep edge, got: %v", e)
	}
	e = g.Edge("edge2")
	if e == nil {
		t.Fatalf("RemoveVertex() expected to keep edge, got: %v", e)
	}
	e = g.Edge("edge3")
	if e == nil {
		t.Fatalf("RemoveVertex() expected to keep edge, got: %v", e)
	}
}

func TestGraph_Vertices(t *testing.T) {
	g := New()
	v := g.Vertices()
	if len(v) > 0 {
		t.Fatalf("Vertices() expected to return empty slice, got: %v", v)
	}
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	v = g.Vertices()
	if len(v) != 2 {
		t.Fatalf("Vertices() expected to return slice with 2 element, got: %v", v)
	}
}

func TestGraph_ForEachVertex(t *testing.T) {
	g := New()
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	visited := map[VertexID]bool{}
	g.ForEachVertex(func(v Vertex) bool {
		visited[v.Id()] = true
		return true
	})
	if len(visited) != 2 {
		t.Fatalf("ForEachVertex() expected to visit 2 vertices, visited: %v", visited)
	}
}

func TestGraph_FOrEachVertex_WithBreak(t *testing.T) {
	g := New()
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	visited := map[VertexID]bool{}
	g.ForEachVertex(func(v Vertex) bool {
		visited[v.Id()] = true
		return false
	})
	if len(visited) != 1 {
		t.Fatalf("ForEachVertex() expected to visit 1 vertices, visited: %v", visited)
	}
}

func TestGraph_AddEdge(t *testing.T) {
	g := New()
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	e, err := g.AddEdge("edge1", "node1", "node2")
	if err != nil {
		t.Fatalf("AddEdge() expected no error, got: %v", err)
	}
	if e == nil {
		t.Fatalf("AddEdge() returned nil edge")
	}
}

func TestGraph_AddEdge_ErrorEdgeAlreadyAdded(t *testing.T) {
	g := New()
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	_, _ = g.AddEdge("edge1", "node1", "node2")
	_, err := g.AddEdge("edge1", "node1", "node2")
	if !errors.Is(err, ErrEdgeAlreadyAdded) {
		t.Fatalf("AddEdge() expected error ErrEdgeAlreadyAdded, got: %v", err)
	}
}

func TestGraph_AddEdgeOnConsistentGraph_ErrorVertexDoesNotExists(t *testing.T) {
	g := New()
	g.EnsureConsistency(true)
	_, err := g.AddEdge("edge1", "node1", "node2")
	if !errors.Is(err, ErrVertexDoesNotExists) {
		t.Fatalf("AddEdge() expected error ErrVertexDoesNotExists, got: %v", err)
	}
	if !strings.Contains(err.Error(), "vertex 'node1' does not exists") {
		t.Fatalf("AddEdge() expected error message to contain 'vertex 'node1' does not exists', got: %v", err)
	}
	_, _ = g.AddVertex("node1")
	_, err = g.AddEdge("edge1", "node1", "node2")
	if !errors.Is(err, ErrVertexDoesNotExists) {
		t.Fatalf("AddEdge() expected error ErrVertexDoesNotExists, got: %v", err)
	}
	if !strings.Contains(err.Error(), "vertex 'node2' does not exists") {
		t.Fatalf("AddEdge() expected error message to contain 'vertex 'node2' does not exists', got: %v", err)
	}
}

func TestGraph_AddEdgeOnNonConsistentGraph(t *testing.T) {
	g := New()
	g.EnsureConsistency(false)
	e, err := g.AddEdge("edge1", "node1", "node2")
	if err != nil {
		t.Fatalf("AddEdge() expected no error, got: %v", err)
	}
	if e == nil {
		t.Fatalf("AddEdge() returned nil edge")
	}
}

func TestGraph_Edge(t *testing.T) {
	g := New()
	e := g.Edge("edge1")
	if e != nil {
		t.Fatalf("Edge() expected to return nil edge, got: %v", e)
	}
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	_, _ = g.AddEdge("edge1", "node1", "node2")
	e = g.Edge("edge1")
	if e == nil {
		t.Fatalf("Edge() expected to return an edge, got: %v", e)
	}
}

func TestGraph_RemoveEdge(t *testing.T) {
	g := New()
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	e, _ := g.AddEdge("edge1", "node1", "node2")
	g.RemoveEdge("edge1")
	e1 := g.Edge("edge1")
	if e1 != nil {
		t.Fatalf("RemoveEdge() expected to remove edge, got: %v", e1)
	}
	if e.Graph() != nil {
		t.Fatalf("RemoveEdge() expected to remove edge from graph, got: %v", e.Graph())
	}
}

func TestGraph_Edges(t *testing.T) {
	g := New()
	e := g.Edges()
	if len(e) > 0 {
		t.Fatalf("Edges() expected to return empty slice, got: %v", e)
	}
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	_, _ = g.AddVertex("node3")
	_, _ = g.AddEdge("edge1", "node1", "node2")
	_, _ = g.AddEdge("edge2", "node2", "node3")
	e = g.Edges()
	if len(e) != 2 {
		t.Fatalf("Edges() expected to return slice with 2 element, got: %v", e)
	}
}

func TestGraph_ForEachEdge(t *testing.T) {
	g := New()
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	_, _ = g.AddEdge("edge1", "node1", "node2")
	_, _ = g.AddEdge("edge2", "node2", "node1")
	visited := map[EdgeID]bool{}
	g.ForEachEdge(func(e Edge) bool {
		visited[e.Id()] = true
		return true
	})
	if len(visited) != 2 {
		t.Fatalf("ForEachEdge() expected to visit 2 edges, visited: %v", visited)
	}
}

func TestGraph_ForEachEdge_WithBreak(t *testing.T) {
	g := New()
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	_, _ = g.AddEdge("edge1", "node1", "node2")
	_, _ = g.AddEdge("edge2", "node2", "node1")
	visited := map[EdgeID]bool{}
	g.ForEachEdge(func(e Edge) bool {
		visited[e.Id()] = true
		return false
	})
	if len(visited) != 1 {
		t.Fatalf("ForEachEdge() expected to visit 1 edges, visited: %v", visited)
	}
}

func TestGraph_EdgesBetween(t *testing.T) {
	g := New()
	e := g.EdgesBetween("node1", "node2")
	if len(e) > 0 {
		t.Fatalf("EdgesBetween() expected to return empty slice, got: %v", e)
	}
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	_, _ = g.AddVertex("node3")
	_, _ = g.AddEdge("edge1", "node1", "node2")
	_, _ = g.AddEdge("edge2", "node2", "node1")
	_, _ = g.AddEdge("edge3", "node1", "node2")
	_, _ = g.AddEdge("edge4", "node1", "node2")
	_, _ = g.AddEdge("edge5", "node2", "node3")
	e = g.EdgesBetween("node1", "node2")
	if len(e) != 3 {
		t.Fatalf("EdgesBetween() expected to return slice with 3 element, got: %v", e)
	}
}

func TestGraph_Consistency(t *testing.T) {
	g := New()
	if !g.EnsuresConsistency() {
		t.Fatalf("EnsuresConsistency() expected to return true on new graph, got: %v", g.EnsuresConsistency())
	}
	g.EnsureConsistency(false)
	if g.EnsuresConsistency() {
		t.Fatalf("EnsuresConsistency() expected to return false, got: %v", g.EnsuresConsistency())
	}
	g.EnsureConsistency(true)
	if !g.EnsuresConsistency() {
		t.Fatalf("EnsuresConsistency() expected to return true, got: %v", g.EnsuresConsistency())
	}
}

func TestGraph_IsConsistent(t *testing.T) {
	g := New()
	if !g.IsConsistent() {
		t.Fatalf("IsConsistent() expected to return true on new graph, got: %v", g.IsConsistent())
	}
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	_, _ = g.AddEdge("edge1", "node1", "node2")
	if !g.IsConsistent() {
		t.Fatalf("IsConsistent() expected to return true, got: %v", g.IsConsistent())
	}
	g.EnsureConsistency(false)
	if !g.IsConsistent() {
		t.Fatalf("IsConsistent() expected to return true, got: %v", g.IsConsistent())
	}
	_, _ = g.AddEdge("edge2", "node1", "node3")
	if g.IsConsistent() {
		t.Fatalf("IsConsistent() expected to return false, got: %v", g.IsConsistent())
	}
	g.RemoveEdge("edge2")
	if !g.IsConsistent() {
		t.Fatalf("IsConsistent() expected to return true, got: %v", g.IsConsistent())
	}
	g.EnsureConsistency(true)
	if !g.IsConsistent() {
		t.Fatalf("IsConsistent() expected to return true, got: %v", g.IsConsistent())
	}
}

func TestGraph_Order(t *testing.T) {
	g := New()
	if g.Order() != 0 {
		t.Fatalf("Order() expected to return 0 on new graph, got: %v", g.Order())
	}
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	if g.Order() != 2 {
		t.Fatalf("Order() expected to return 2, got: %v", g.Order())
	}
}

func TestGraph_Size(t *testing.T) {
	g := New()
	if g.Size() != 0 {
		t.Fatalf("Size() expected to return 0 on new graph, got: %v", g.Size())
	}
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	_, _ = g.AddEdge("edge1", "node1", "node2")
	_, _ = g.AddEdge("edge2", "node1", "node2")
	if g.Size() != 2 {
		t.Fatalf("Size() expected to return 2, got: %v", g.Size())
	}
}

func TestGraph_Degree(t *testing.T) {
	g := New()
	if g.Degree() != 0 {
		t.Fatalf("Degree() expected to return 0 on new graph, got: %v", g.Degree())
	}
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	_, _ = g.AddVertex("node3")
	_, _ = g.AddEdge("edge1", "node1", "node2")
	_, _ = g.AddEdge("edge2", "node2", "node1")
	_, _ = g.AddEdge("edge3", "node1", "node3")
	if g.Degree() != 3 {
		t.Fatalf("Degree() expected to return 3, got: %v", g.Degree())
	}
	_, _ = g.AddEdge("edge4", "node1", "node1")
	if g.Degree() != 5 {
		t.Fatalf("Degree() expected to return 4, got: %v", g.Degree())
	}
	g.RemoveVertex("node1")
	if g.Degree() != 0 {
		t.Fatalf("Degree() expected to return 1, got: %v", g.Degree())
	}
}

func TestGraph_Clone(t *testing.T) {
	g := New()
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	_, _ = g.AddVertex("node3")
	_, _ = g.AddEdge("edge1", "node1", "node2")
	_, _ = g.AddEdge("edge2", "node2", "node3")
	_, _ = g.AddEdge("edge3", "node3", "node3")
	clone := g.Clone()
	if clone == nil {
		t.Fatalf("Clone() returned nil graph")
	}
	if len(clone.Vertices()) != len(g.Vertices()) {
		t.Fatalf("Clone() expected to have same number of vertices, got: %v", len(clone.Vertices()))
	}
	if len(clone.Edges()) != len(g.Edges()) {
		t.Fatalf("Clone() expected to have same number of edges, got: %v", len(clone.Edges()))
	}
	_, _ = clone.AddVertex("node4")
	if g.Vertex("node4") != nil {
		t.Fatalf("Clone() expected to be a deep copy, got: %v", g.Vertex("node4"))
	}
	_, _ = clone.AddEdge("edge4", "node1", "node4")
	if g.Edge("edge4") != nil {
		t.Fatalf("Clone() expected to be a deep copy, got: %v", g.Edge("edge4"))
	}
}
