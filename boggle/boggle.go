package boggle

type Game struct {
	letters [][]rune
}

func New(rows []string) *Game {
	row_cnt := len(rows)
	if row_cnt < 1 {
		return nil
	}

	g := new(Game)
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

	return g
}
