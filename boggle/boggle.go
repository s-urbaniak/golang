package boggle

type Game struct {
	size    BoardSize
	adj     [][]Vertex
	letters []rune
}

type BoardSize struct {
	m, n int
}

type Vertex int

type DfsCb func(Vertex)

func dfs(cb DfsCb, adj [][]Vertex, u Vertex) {
	visited := make(map[Vertex]bool)
	dfsVisit(cb, visited, adj, u)
}

func dfsVisit(cb DfsCb, visited map[Vertex]bool, adj [][]Vertex, u Vertex) {
	cb(u)
	visited[u] = true

	for _, v := range adj[u] {
		if !visited[v] {
			dfsVisit(cb, visited, adj, v)
		}
	}
}

func NewVertex(i int, s *BoardSize) Vertex {
	if i >= 0 && i < s.m*s.n {
		return Vertex(i)
	} else {
		return -1
	}
}

func NewVertexFromCoords(x, y int, s *BoardSize) Vertex {
	if x >= 0 && y >= 0 && x < s.n && y < s.m {
		return Vertex(y*s.m + x)
	} else {
		return -1
	}
}

func (v *Vertex) IterNeighbours(s *BoardSize) chan Vertex {
	ch := make(chan Vertex)

	vx := int(*v) % s.n
	vy := (int(*v) - vx) / s.m

	go func() {
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				n := NewVertexFromCoords(vx+x, vy+y, s)
				if n >= 0 && n != *v {
					ch <- n
				}
			}
		}

		close(ch)
	}()

	return ch
}

func (g *Game) Init(words []string) *Game {
	rows := len(words)
	if rows < 1 {
		return nil
	}

	cols := 0
	for m := 0; m < rows; m++ {
		row := []rune(words[m])

		if cols == 0 {
			cols = len(row)
			size := rows * cols

			g.letters = make([]rune, size)
			g.adj = make([][]Vertex, size)
			g.size = BoardSize{rows, cols}
		} else if cols != len(row) {
			return nil
		}

		for n := 0; n < cols; n++ {
			i := m*rows + n
			g.letters[i] = row[n]

			g.adj[i] = make([]Vertex, 0)
			v := Vertex(i)
			for a := range v.IterNeighbours(&g.size) {
				g.adj[i] = append(g.adj[i], a)
			}
		}
	}

	return g
}

func NewGame(rows []string) *Game {
	return new(Game).Init(rows)
}
