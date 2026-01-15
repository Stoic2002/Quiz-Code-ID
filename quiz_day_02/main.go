package main

import "fmt"

func main() {
	// displayMatrix(matrix1([5][5]int{}))
	// displayMatrix(matrix2([5][5]int{}))
	// displayMatrix(matrix3([7][7]int{}))
	// displayMatrix(matrix4([7][7]int{}))
	correctTheAnswer()
}

func displayMatrix(matrix [8][8]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Println()
	}
}

// nomor 6
func matrix1(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == j {
				matrix[i][j] = i + 1
			} else if i <= j {
				matrix[i][j] = 10
			} else if i >= j {
				matrix[i][j] = 20
			}
		}
	}
	return matrix
}

// nomor 7
func matrix2(matrix [5][5]int) [5][5]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == j {
				matrix[i][j] = len(matrix) - i
			} else if i <= j {
				matrix[i][j] = 20
			} else if i >= j {
				matrix[i][j] = 10
			}
		}
	}
	return matrix
}

// nomor 8
func matrix3(matrix [7][7]int) [7][7]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 {
				matrix[i][j] = j - i
			} else if j == 0 {
				matrix[i][j] = i
			} else if i == len(matrix)-1 {
				matrix[i][j] = len(matrix) + j - 1
			} else if j == len(matrix[i])-1 {
				matrix[i][j] = len(matrix[i]) + i - 1
			}
		}
	}
	return matrix
}

// nomor 9
func matrix4(matrix [7][7]int) [8][8]int {
	res := [8][8]int{}

	for i := 0; i < len(matrix); i++ {
		rowSum := 0

		for j := 0; j < len(matrix[i]); j++ {

			val := i + j

			res[i][j] = val

			rowSum += val

			res[len(res)-1][j] += val

		}
		res[i][len(res)-1] = rowSum

		res[len(res)-1][len(res)-1] += rowSum
	}
	return res
}

// nomor 10
func correctTheAnswer() {

	studentAnswerList := [][]string{
		{"A", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"D", "B", "A", "B", "C", "A", "E", "E", "A", "D"},
		{"E", "D", "D", "A", "C", "B", "E", "E", "A", "D"},
		{"C", "B", "A", "E", "D", "C", "E", "E", "A", "D"},
		{"A", "B", "D", "C", "C", "D", "E", "E", "A", "D"},
		{"B", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
		{"B", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"E", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
	}

	correctAnswer := []string{"D", "B", "D", "C", "C", "D", "A", "E", "A", "D"}

	for i, answers := range studentAnswerList {
		correct := 0
		for j, answer := range answers {
			if answer == correctAnswer[j] {
				correct++
			}
		}
		fmt.Println("Jawaban Siswa", i, "yang benar :", correct)
	}

}
