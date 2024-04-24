package main

func DetectMagicSquare(square [][]int) bool {
	size := len(square)
	row_sum := make([]int, size)
	col_sum := make([]int, size)
	diag_sum := 0
	rev_diag_sum := 0
	sums := make([]int, 2*size+2)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			row_sum[i] += square[i][j]
			col_sum[j] += square[i][j]
			if i == j {
				diag_sum += square[i][j]
			}
			if i+j == size-1 {
				rev_diag_sum += square[i][j]
			}
		}
	}
	sums[0] = diag_sum
	sums[1] = rev_diag_sum
	for i := 2; i < size+2; i++ {
		sums[i] = row_sum[i-2]
	}
	for i := size + 2; i < 2*size+2; i++ {
		sums[i] = col_sum[i-size-2]
	}
	for i := 1; i < 2*size+2; i++ {
		if sums[i] != sums[0] {
			return false
		}
	}
	return true
}
