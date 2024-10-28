package mgraph

import "testing"

func TestEdge_Id(t *testing.T) {
	g := New()
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	e, _ := g.AddEdge("edge1", "node1", "node2")
	if e.Id() != "edge1" {
		t.Errorf("Id() expected 'edge1', got %s", e.Id())
	}
}

func TestEdge_FromAndTo(t *testing.T) {
	g := New()
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	e, _ := g.AddEdge("edge1", "node1", "node2")
	if e.From() != "node1" {
		t.Errorf("From() expected 'node1', got %s", e.From())
	}
	if e.To() != "node2" {
		t.Errorf("To() expected 'node2', got %s", e.To())
	}
}

func TestEdge_Data(t *testing.T) {
	g := New()
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	e, _ := g.AddEdge("edge1", "node1", "node2")
	if e.Data() != nil {
		t.Errorf("Data() expected nil, got %v", e.Data())
	}
	e.StoreData("data")
	if e.Data() != "data" {
		t.Errorf("Data() expected 'data', got %s", e.Data())
	}
}

func TestEdge_TailAndHeadAndEndpoints_WithConsistentGraph(t *testing.T) {
	g := New()
	v1, _ := g.AddVertex("node1")
	v2, _ := g.AddVertex("node2")
	e, _ := g.AddEdge("edge1", "node1", "node2")
	if e.Tail() != v1 {
		t.Errorf("Tail() expected %v, got %v", v1, e.Tail())
	}
	if e.Head() != v2 {
		t.Errorf("Head() expected %v, got %v", v2, e.Head())
	}
	if e.Endpoints() != [2]Vertex{v1, v2} {
		t.Errorf("Endpoints() expected [%v %v], got %v", v1, v2, e.Endpoints())
	}
}

func TestEdge_TailAndHeadEndpoints_WithInconsistentGraph(t *testing.T) {
	g := New()
	g.EnsureConsistency(false)
	e, _ := g.AddEdge("edge1", "node1", "node2")
	if e.Tail() != nil {
		t.Errorf("Tail() expected nil, got %v", e.Tail())
	}
	if e.Head() != nil {
		t.Errorf("Head() expected nil, got %v", e.Head())
	}
	if e.Endpoints() != [2]Vertex{nil, nil} {
		t.Errorf("Endpoints() expected [nil nil], got %v", e.Endpoints())
	}
}

func TestEdge_IsIncident(t *testing.T) {
	g := New()
	_, _ = g.AddVertex("node1")
	v2, _ := g.AddVertex("node2")
	v3, _ := g.AddVertex("node3")
	e, _ := g.AddEdge("edge1", "node1", "node2")
	if !e.IsIncident(v2) {
		t.Errorf("IsIncident() expected true, got false")
	}
	if e.IsIncident(v3) {
		t.Errorf("IsIncident() expected false, got true")
	}
}

func TestEdge_IsInverted(t *testing.T) {
	g := New()
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	_, _ = g.AddVertex("node3")
	e1, _ := g.AddEdge("edge1", "node1", "node2")
	e2, _ := g.AddEdge("edge2", "node2", "node1")
	e3, _ := g.AddEdge("edge3", "node2", "node3")
	if !e1.IsInverted(e2) {
		t.Errorf("IsInverted() expected true, got false")
	}
	if !e2.IsInverted(e1) {
		t.Errorf("IsInverted() expected true, got false")
	}
	if e1.IsInverted(e3) {
		t.Errorf("IsInverted() expected false, got true")
	}
	if e3.IsInverted(e1) {
		t.Errorf("IsInverted() expected false, got true")
	}
}

func TestEdge_IsParallel(t *testing.T) {
	g := New()
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	_, _ = g.AddVertex("node3")
	e1, _ := g.AddEdge("edge1", "node1", "node2")
	e2, _ := g.AddEdge("edge2", "node1", "node2")
	e3, _ := g.AddEdge("edge3", "node2", "node3")
	if !e1.IsParallel(e2) {
		t.Errorf("IsParallel() expected true, got false")
	}
	if !e2.IsParallel(e1) {
		t.Errorf("IsParallel() expected true, got false")
	}
	if e1.IsParallel(e3) {
		t.Errorf("IsParallel() expected false, got true")
	}
	if e3.IsParallel(e1) {
		t.Errorf("IsParallel() expected false, got true")
	}
}

func TestEdge_IsLoop(t *testing.T) {
	g := New()
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	e1, _ := g.AddEdge("edge1", "node1", "node2")
	e2, _ := g.AddEdge("edge2", "node1", "node1")
	if e1.IsLoop() {
		t.Errorf("IsLoop() expected false, got true")
	}
	if !e2.IsLoop() {
		t.Errorf("IsLoop() expected true, got false")
	}
}

func TestEdge_Graph(t *testing.T) {
	g := New()
	_, _ = g.AddVertex("node1")
	_, _ = g.AddVertex("node2")
	e, _ := g.AddEdge("edge1", "node1", "node2")
	if e.Graph() != g {
		t.Errorf("Graph() expected %v, got %v", g, e.Graph())
	}
}
