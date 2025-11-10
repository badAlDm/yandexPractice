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

	var nums [3]int

	for i := 0; i < 3; i++ {
		nums[i] = readInt(reader)

	}

	_, err := writer.WriteString(isTriangle(nums))
	if err != nil {
		panic(err)
	}

	writer.WriteByte('\n')
}

func readInt(r *bufio.Reader) int {
	line, _ := r.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	return n
}

func isTriangle(sides [3]int) string {
	if sides[0]+sides[1] > sides[2] && sides[1]+sides[2] > sides[0] && sides[0]+sides[2] > sides[1] {
		return "YES"
	}

	return "NO"
}
