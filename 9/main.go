package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)

	line, _ := reader.ReadString('\n')

	nmk := lineToIntArr(line)
	if len(nmk) != 3 {
		fmt.Println("error: bad input")

		return
	}

	mA := make([][]int, nmk[0])
	mB := make([][]int, nmk[1])

	for i := 0; i < nmk[0]; i++ {
		line, _ := reader.ReadString('\n')
		mA[i] = lineToIntArr(line)
	}

	for i := 0; i < nmk[1]; i++ {
		line, _ := reader.ReadString('\n')
		mB[i] = lineToIntArr(line)
	}

	var sb strings.Builder
	for _, row := range Transp(MultiplyM(mA, mB)) {
		for j, val := range row {
			if j > 0 {
				sb.WriteByte(' ')
			}
			fmt.Fprintf(&sb, "%d", val)
		}
		sb.WriteByte('\n')
	}

	fmt.Println(sb.String())

}

func lineToIntArr(line string) []int {
	line = strings.TrimSpace(line)
	result := make([]int, 0)
	for _, num := range strings.Split(line, " ") {
		num, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}

		result = append(result, num)
	}

	return result
}

func Transp(m [][]int) [][]int {
	result := make([][]int, len(m[0]))
	for i := range result {
		result[i] = make([]int, len(m))
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			result[j][i] = m[i][j]
		}
	}

	return result
}

func MultiplyM(A, B [][]int) [][]int {
	if len(A[0]) != len(B) {
		panic(errors.New("bad matrix"))
	}

	C := make([][]int, len(A))
	for i := range C {
		C[i] = make([]int, len(B[0]))
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B[0]); j++ {
			for k := 0; k < len(A[0]); k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return C
}
