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

	b.Dfs(Vertex(4), "efi", func(v Vertex) {
		t.Log(v)
	})
}
