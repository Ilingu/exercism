package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	matrix  [][]int
	rows    int
	columns int
}

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")

	fomattedRows, noOfCols := [][]int{}, -1
	for _, row := range rows {
		points, intPoints := strings.Split(strings.TrimSpace(row), " "), []int{}
		for _, point := range points {
			intPoint, err := strconv.Atoi(point)
			if err != nil {
				return nil, errors.New("invalid number")
			}

			intPoints = append(intPoints, intPoint)
		}

		if noOfCols == -1 {
			noOfCols = len(points)
		}
		if len(points) != noOfCols {
			return nil, errors.New("non continuous columns")
		}

		fomattedRows = append(fomattedRows, intPoints)
	}

	if noOfCols == -1 {
		return nil, errors.New("no columns")
	}

	matrix := append([][]int{}, fomattedRows...)
	return &Matrix{matrix: matrix, rows: len(rows), columns: noOfCols}, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	columns := make([][]int, m.columns)
	for _, row := range m.matrix {
		for i := 0; i < m.columns; i++ {
			columns[i] = append(columns[i], row[i])
		}
	}
	return columns
}

func (m *Matrix) Rows() [][]int {
	rows := make([][]int, m.rows)
	for i, row := range m.matrix {
		rows[i] = append(rows[i], row...)
	}

	return rows
}

func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row > m.rows-1 || col > m.columns-1 {
		return false
	}

	m.matrix[row][col] = val
	return true
}
