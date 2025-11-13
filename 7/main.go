package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)

	words := make(map[string]struct{})
	for scan.Scan() {
		words[scan.Text()] = struct{}{}
	}

	fmt.Println(len(words))
}
