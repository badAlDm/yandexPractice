package main

import (
	"bufio"
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

	nums := lineToIntArr(str)

	_, err = writer.WriteString(yesOrNo(nums))
	if err != nil {
		panic(err)
	}

	writer.WriteByte('\n')

}

func yesOrNo(nums []int) string {

	for i, n := range nums {
		if i == 0 {
			continue
		}
		if nums[i-1] >= n {
			return "NO"
		}
	}
	return "YES"
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
