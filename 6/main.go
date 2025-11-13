package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var inpt [2]string

	for i := 0; i < 2; i++ {
		inpt[i] = readStr(reader)

	}
	existButtons := lineToIntArr(inpt[0])
	wantNum := inpt[1]

	_, err := writer.WriteString(fmt.Sprintf("%d", wantsB(existButtons, wantNum)))
	if err != nil {
		panic(err)
	}

	writer.WriteByte('\n')
}

func readStr(r *bufio.Reader) string {
	line, _ := r.ReadString('\n')
	return strings.TrimSpace(line)
}
func lineToIntArr(line string) map[string]struct{} {
	result := make(map[string]struct{})
	for _, num := range strings.Split(line, " ") {
		result[num] = struct{}{}
	}

	return result
}

func wantsB(buttons map[string]struct{}, number string) int {
	res := make(map[string]struct{})
	for _, s := range strings.Split(number, "") {
		_, exist := buttons[s]
		if !exist {
			res[s] = struct{}{}
		}
	}

	return len(res)
}
