package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	in := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)
	for {
		if ok := scanner.Scan(); !ok {
			break
		} else {
			word := scanner.Text()
			count[word]++
		}
	}
	fmt.Println("Word\tCount")
	for k, v := range count {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
