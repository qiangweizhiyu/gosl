// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package num

import "github.com/cpmech/gosl/chk"

// scalar function
type Fun0 func(float64) float64

// Simpson integrates function f from x=a to x=b using n subintervals (n must be even)
func Simpson(f Fun0, a, b float64, n int) (float64, error) {
	if n < 2 || n%2 > 0 {
		return 0, chk.Err("number of subintervas should be even (n=%d)", n)
	}

	h := (b - a) / float64(n)
	sum := f(a) + f(b)

	x := a
	for i := 1; i < n; i++ {
		x += h
		if i%2 == 1 { // i is odd
			sum += 4 * f(x)
		} else { // i is even
			sum += 2 * f(x)
		}
	}

	sum = sum * h / 3

	return sum, nil
}

// Simps2D computes a double integral over the x-y plane (Simpson's rule); thus resulting on
// the volume defined between the function f(x,y) and the plane ortogonal to z
func Simps2D(dx, dy float64, f [][]float64) (V float64) {

	// check
	if len(f) < 2 {
		chk.Panic("len(f)=%d is incorrect; it must be at least 2", len(f))
	}
	m, n := len(f), len(f[0])

	// corners
	V = f[0][0] + f[m-1][0] + f[0][n-1] + f[m-1][n-1]

	// top/bottom: 4
	for j := 1; j < n-1; j += 2 {
		V += 4.0 * (f[0][j] + f[m-1][j])
	}

	// top/bottom: 2
	for j := 2; j < n-1; j += 2 {
		V += 2.0 * (f[0][j] + f[m-1][j])
	}

	// left/right: 4
	for i := 1; i < m-1; i += 2 {
		V += 4.0 * (f[i][0] + f[i][n-1])

		// centre: 16
		for j := 1; j < n-1; j += 2 {
			V += 16.0 * f[i][j]
		}

		// centre: 8a
		for j := 2; j < n-1; j += 2 {
			V += 8.0 * f[i][j]
		}
	}

	// left/right: 2
	for i := 2; i < m-1; i += 2 {
		V += 2.0 * (f[i][0] + f[i][n-1])

		// centre: 4
		for j := 2; j < n-1; j += 2 {
			V += 4.0 * f[i][j]
		}

		// centre: 8b
		for j := 1; j < n-1; j += 2 {
			V += 8.0 * f[i][j]
		}
	}

	// final result
	V *= dx * dy / 9.0
	return
}
