package main

import (
	"errors"
	"fmt"
)

const N = 5

func main() {

	matrix := make([][]int, N)
	for i := 0; i < N; i++ {
		matrix[i] = make([]int, N)
	}
	matrix, err := generateMatrix(matrix, 0, 0, N)
	if err != nil {
		fmt.Println(err)
	}
	//squiralMatrix := [][]int{}
	for i := 0; i < N; i++ {
		subLine := make([]int, 0, N)
		for j := 0; j < N; j++ {
			subLine = append(subLine, matrix[i][j])
		}
		fmt.Printf("%v\n", subLine)
	}
}
func generateMatrix(matrix [][]int, x, y, N int) ([][]int, error) {

	if matrix == nil || x < 0 || y < 0 {
		return nil, errors.New("input error")
	}

	val := 1
	direction := 1
	for {
		if val > N*N-1 {
			matrix[x][y] = val
			break
		}

		switch direction {
		case 1: // ->
			if y+1 < N && matrix[x][y+1] == 0 {
				matrix[x][y] = val
				val++
				y++
			} else {
				direction = 2
			}
		case 2: // |
			if x+1 < N && matrix[x+1][y] == 0 {
				matrix[x][y] = val
				val++
				x++
			} else {
				direction = 3
			}
		case 3: // <-
			if y-1 >= 0 && matrix[x][y-1] == 0 {
				matrix[x][y] = val
				val++
				y--
			} else {
				direction = 4
			}
		case 4: // ^
			if x-1 >= 0 && matrix[x-1][y] == 0 {
				matrix[x][y] = val
				val++
				x--
			} else {
				direction = 1
			}
		default:
		}
	}
	return matrix, nil
}
