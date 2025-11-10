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

	nums := lineToIntArr(str)
	if len(nums) != 2 {
		fmt.Println("error: bad input")

		return
	}

	writer.WriteString(fmt.Sprintf("%d", nums[0]+nums[1]))
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
