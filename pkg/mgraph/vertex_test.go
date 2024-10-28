package mgraph

import "testing"

func TestVertex_Id(t *testing.T) {
	g := New()
	v, _ := g.AddVertex("node1")
	if v.Id() != "node1" {
		t.Errorf("Id() expected 'node1', got %s", v.Id())
	}
}

func TestVertex_Data(t *testing.T) {
	g := New()
	v, _ := g.AddVertex("node1")
	if v.Data() != nil {
		t.Errorf("Data() expected nil, got %v", v.Data())
	}
	v.StoreData("data")
	if v.Data() != "data" {
		t.Errorf("Data() expected 'data', got %s", v.Data())
	}
}

func TestVertex_Graph(t *testing.T) {
	g := New()
	v, _ := g.AddVertex("node1")
	if v.Graph() != g {
		t.Errorf("Graph() expected %v, got %v", g, v.Graph())
	}
}

func TestVertex_IncomingAndOutgoingAndEdges(t *testing.T) {
	g := New()
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	_, _ = g.AddVertex("node3")
	_, _ = g.AddEdge("edge1", "node1", "node2")
	_, _ = g.AddEdge("edge2", "node1", "node3")
	_, _ = g.AddEdge("edge3", "node3", "node1")
	v := g.Vertex("node1")
	incoming := v.Incoming()
	if len(incoming) != 1 {
		t.Errorf("Incoming() expected 1, got %d", len(incoming))
	}
	outgoing := v.Outgoing()
	if len(outgoing) != 2 {
		t.Errorf("Outgoing() expected 2, got %d", len(outgoing))
	}
	edges := v.Edges()
	if len(edges) != 3 {
		t.Errorf("Edges() expected 3, got %d", len(edges))
	}
}

func TestVertex_BelongsTo(t *testing.T) {
	g := New()
	v, _ := g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	e, _ := g.AddEdge("edge1", "node1", "node2")
	e2, _ := g.AddEdge("edge2", "node2", "node2")
	if !v.BelongsTo(e) {
		t.Errorf("BelongsTo() expected true, got false")
	}
	if v.BelongsTo(e2) {
		t.Errorf("BelongsTo() expected false, got true")
	}
}

func TestVertex_Degree(t *testing.T) {
	g := New()
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	_, _ = g.AddVertex("node3")
	_, _ = g.AddEdge("edge1", "node1", "node2")
	_, _ = g.AddEdge("edge2", "node1", "node3")
	_, _ = g.AddEdge("edge3", "node3", "node1")
	v := g.Vertex("node1")
	if v.Degree() != 3 {
		t.Errorf("Degree() expected 3, got %d", v.Degree())
	}
}
