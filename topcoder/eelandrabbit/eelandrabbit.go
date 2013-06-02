package eelandrabbit

type EeelAndRabbit interface {
	GrabEels(l, t []int) (m int)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func CatchOne(l, t []int) (m int) {
	if len(l) != len(t) {
		return 0
	}

	m = 0

	for a := 0; a < len(l); a++ {
		grabbed := 0
		head_a := t[a]

		for b := 0; b < len(l); b++ {
			head_b := t[b]
			tail_b := t[b] + l[b]

			if head_b <= head_a && tail_b >= head_a {
				grabbed++
			}
		}

		m = max(m, grabbed)
	}

	return
}

func CatchTwo(l, t []int) (m int) {
	if len(l) != len(t) {
		return 0
	}

	m = 0

	for a := 0; a < len(l); a++ {
		for b := 0; b < len(l); b++ {
			grabbed := 0
			head_a := t[a]
			head_b := t[b]

			for c := 0; c < len(l); c++ {
				head_c := t[c]
				tail_c := t[c] + l[c]

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
