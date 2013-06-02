package eelandrabbit

type EelAndRabbit struct {
	l, t []int
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (e *EelAndRabbit) CatchOne() (m int) {
	if len(e.l) != len(e.t) {
		return 0
	}

	m = 0

	for a := 0; a < len(e.l); a++ {
		grabbed := 0
		head_a := e.t[a]

		for b := 0; b < len(e.l); b++ {
			head_b := e.t[b]
			tail_b := e.t[b] + e.l[b]

			if head_b <= head_a && tail_b >= head_a {
				grabbed++
			}
		}

		m = max(m, grabbed)
	}

	return
}

func (e *EelAndRabbit) CatchTwo() (m int) {
	if len(e.l) != len(e.t) {
		return 0
	}

	m = 0

	for a := 0; a < len(e.l); a++ {
		for b := 0; b < len(e.l); b++ {
			grabbed := 0
			head_a := e.t[a]
			head_b := e.t[b]

			for c := 0; c < len(e.l); c++ {
				head_c := e.t[c]
				tail_c := e.t[c] + e.l[c]

				if (head_c <= head_a && tail_c >= head_a) ||
					(head_c <= head_b && tail_c >= head_b) {
					grabbed++
				}
			}

			m = max(m, grabbed)
		}
	}

	return
}
