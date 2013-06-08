package boggle

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	b := NewGame([]string{"abc", "def", "ghi"})
	t.Log(b, len(b.adj[0]))

	if b == nil {
		t.Error("expecting non nil value")
	}

	found := b.Dfs(Vertex(0), "abefefe", func(v Vertex) {
		t.Log(string(b.letters[v]))
	})

	if !found {
		t.Error("not found")
	}
}
