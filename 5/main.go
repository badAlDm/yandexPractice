package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	str := strings.TrimSpace(line)
	nums := lineToIntArr(str)
	fmt.Println(nums)
	if len(nums) != 3 {
		panic(errors.New("bad input"))
	}

	findSq(nums[0], nums[1], nums[2])
	// _, err = writer.WriteString(fmt.Sprintf("%d", len(res)))
	// if err != nil {
	// 	panic(err)
	// }

	// writer.WriteByte('\n')

	// slices.Sort(res)

	// _, err = writer.WriteString(fmt.Sprintf("%s", strings.Join(res, " ")))
	// if err != nil {
	// 	panic(err)
	// }

	// writer.WriteByte('\n')
}

func lineToIntArr(line string) []float64 {
	result := make([]float64, 0)
	for _, num := range strings.Split(line, " ") {
		num, err := strconv.ParseFloat(num, 64)

		if err != nil {
			panic(err)
		}

		result = append(result, num)
	}

	return result
}

func findSq(a, b, c float64) {
	D := math.Pow(b, 2) - 4*a*c
	switch {
	case D > 0:
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)

		res := []float64{x1, x2}

		slices.Sort(res)

		fmt.Println(2)
		fmt.Printf("%.6f %.6f\n", res[0], res[1])

		return
	case D == 0:
		x := -b / (2 * a)

		fmt.Println(1)
		fmt.Printf("%.6f\n", x)

		return
	default:
		fmt.Println(0)

		return
	}

}
