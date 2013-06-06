package boggle

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	b := NewGame([]string{"abc", "def", "ghi"})
	t.Log(b, len(b.adj[0]))

	// abc
	// def
	// ghi
	dfs(func(v Vertex) {
		t.Log(v)
	}, b.adj, Vertex(8))

	if b == nil {
		t.Error("expecting non nil value")
	}
}
