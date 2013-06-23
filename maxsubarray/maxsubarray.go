// Package maxsubarray provides functions for calculating the max subarray value
package maxsubarray

func maxcross(A []int, a, m, b int) (l, r, sum int) {
	lmax, rmax, sum := 0, 0, 0

	// invariant: i is in the left part of the input array (a <= i <= m)
	for i := m; i >= a; i-- {
		sum += A[i]

		if sum > lmax {
			lmax = sum
			l = i
		}
	}

	sum = 0
	// invariant: i is in the right part of the input array (m < i < b)
	for i := m + 1; i < b; i++ {
		sum += A[i]

		if sum > rmax {
			rmax = sum
			r = i
		}
	}

	return l, r, lmax + rmax
}

func maxsubarray(A []int, a, b int) (l, r, sum int) {
	if a == b-1 {
		return a, a, A[a]
	} else {
		m := a + (b-1-a)/2

		l1, r1, sum1 := maxsubarray(A, a, m+1)
		l2, r2, sum2 := maxsubarray(A, m+1, b)
		lc, rc, sumc := maxcross(A, a, m, b)

		if sum1 >= sum2 && sum1 >= sumc {
			return l1, r1, sum1
		} else if sum2 >= sum1 && sum2 >= sumc {
			return l2, r2, sum2
		} else {
			return lc, rc, sumc
		}
	}
}
