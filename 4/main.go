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
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	str := strings.TrimSpace(line)

	res := tagCounter(lineToInt(str))

	_, err = writer.WriteString(fmt.Sprintf("%d", res))
	if err != nil {
		panic(err)
	}

	writer.WriteByte('\n')
}

func lineToInt(line string) int {
	num, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	if num < 1 || num > 35 {
		panic(errors.New("bad input"))
	}

	return num
}

func tagCounter(size int) int {
	result := 0
	count := make([]int, size)
	for i := range size {
		switch i {
		case 0, 1:
			count[i] = 1
			result++
		default:
			count[i] = count[i-1] + count[i-2]
			result += count[i]
		}
	}

	return result
}
