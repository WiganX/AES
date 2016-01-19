package gfmatrix

import (
	"github.com/OpenWhiteBox/AES/primitives/number"
)

// gaussJordan reduces the matrix according to the Gauss-Jordan Method.  Returns the augment matrix, the transformed
// matrix, and the failing pivot (-1 if none).
func (e Matrix) gaussJordan() (aug, f Matrix, frees []int) {
	out, in := e.Size()

	aug = GenerateIdentity(in)
	for x := in; x < out; x++ {
		aug = append(aug, Row(make([]number.ByteFieldElem, in)))
	}

	f = Matrix(make([]Row, out)) // Duplicate e away so we don't mutate it.
	copy(f, e)

	row, col := 0, 0

	for row < out && col < in {
		// Find a non-zero element to move into the pivot position.
		i := f.FindPivot(row, col)
		if i == -1 { // Failed to find a pivot.
			frees = append(frees, col)
			col++

			continue
		}

		// Move it into position.
		f[row], f[i] = f[i], f[row]
		aug[row], aug[i] = aug[i], aug[row]

		// Normalize the entry in this rows (pivot)th position.
		correction := f[row][col].Invert()
		f[row] = f[row].ScalarMul(correction)
		aug[row] = aug[row].ScalarMul(correction)

		// Cancel out the (row)th position for every row above and below it.
		for j, _ := range f {
			if j != row && !f[j][col].IsZero() {
				aug[j] = aug[j].Add(aug[row].ScalarMul(f[j][col]))
				f[j] = f[j].Add(f[row].ScalarMul(f[j][col]))
			}
		}

		row++
		col++
	}

	// Add the rest of the free variables for completion.
	for x := out; x < in; x++ {
		frees = append(frees, x)
	}

	return
}