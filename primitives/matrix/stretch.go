package matrix

// RightStretch returns the matrix of right matrix multiplication by a given matrix.
func (e Matrix) RightStretch() Matrix {
	n, m := e.Size()
	nm := n * m

	out := make([]Row, nm)

	for i := 0; i < nm; i++ {
		out[i] = make([]byte, nm/8)
		p, q := i/n, i%n

		for j := 0; j < m; j++ {
			out[i].SetBit(j*m+q, e[p].GetBit(j) == 1)
		}
	}

	return out
}

// LeftStretch returns the matrix of left matrix multiplication by a given matrix.
func (e Matrix) LeftStretch() Matrix {
	n, m := e.Size()
	nm := n * m

	out := make([]Row, nm)

	for i := 0; i < nm; i++ {
		out[i] = make([]byte, nm/8)
		p, q := i/n, i%n

		for j := 0; j < m; j++ {
			out[i].SetBit(j+m*p, e[j].GetBit(q) == 1)
		}
	}

	return out
}
