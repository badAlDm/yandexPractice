package main

import (
	"bufio"
	"fmt"
	"os"
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

	res := neighbours(lineToIntArr(str))

	_, err = writer.WriteString(fmt.Sprintf("%d", res))
	if err != nil {
		panic(err)
	}

	writer.WriteByte('\n')

}

func lineToIntArr(line string) []int {
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

func neighbours(nums []int) int {
	result := 0
	for i, num := range nums {
		switch i {
		case 0:
			continue
		case len(nums) - 1:
			break
		default:
			if num > nums[i-1] && num > nums[i+1] {
				result++
			}
		}
	}

	return result
}
