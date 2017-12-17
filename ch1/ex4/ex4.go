package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	for _, path := range os.Args[1:] {
		err := countLines(path, counts)
		if err != nil {
			fmt.Println("Error counting lines for file", path, err)
			os.Exit(-1)
		}
	}
	for line, paths := range counts {
		if len(paths) > 1 {
			fmt.Printf("%s: ", line)
			for k, v := range paths {
				fmt.Printf("%s (%d) ", k, v)
			}
			fmt.Println()
		}
	}
}

func countLines(path string, counts map[string]map[string]int) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	input := bufio.NewScanner(f)
	for input.Scan() {
		k := input.Text()
		if _, ok := counts[k]; !ok {
			counts[k] = make(map[string]int)
		}
		counts[k][path]++
	}
	return input.Err()
}
