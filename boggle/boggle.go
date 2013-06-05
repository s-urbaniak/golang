package boggle

type Game struct {
	letters [][]rune
	size    BoardSize
}

type BoardSize struct {
	m, n int
}

type Pos struct {
	x, y int
}

func (p *Pos) IsValid(s *BoardSize) bool {
	return p.x >= 0 && p.y >= 0 && p.x < s.n && p.y < s.m
}

func (g *Game) IterNeighbours(p Pos) chan Pos {
	ch := make(chan Pos)

	if !p.IsValid(&g.size) {
		close(ch)
	} else {
		go func() {
			for y := p.y - 1; y <= p.y+1; y++ {
				for x := p.x - 1; x <= p.x+1; x++ {
					n := Pos{x, y}
					if n.IsValid(&g.size) && n != p {
						ch <- n
					}
				}
			}

			close(ch)
		}()
	}

	return ch
}

func (g *Game) Init(rows []string) *Game {
	row_cnt := len(rows)
	if row_cnt < 1 {
		return nil
	}

	g.letters = make([][]rune, row_cnt)

	col_cnt := 0
	for m := 0; m < row_cnt; m++ {
		row := []rune(rows[m])
		g.letters[m] = row

		if col_cnt == 0 {
			col_cnt = len(row)
		} else if col_cnt != len(row) {
			return nil
		}
	}

	g.size = BoardSize{row_cnt, col_cnt}
	return g
}

func New(rows []string) *Game {
	return new(Game).Init(rows)
}
