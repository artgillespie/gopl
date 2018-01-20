package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

const (
	LETTER = iota
	CONTROL
	GRAPHIC
	DIGIT
	MARK
	SYMBOL
	SPACE
	INVALID
)

func main() {
	count := make(map[int]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "counts: %v\n", count)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			count[INVALID]++
			continue
		}
		if unicode.IsLetter(r) {
			count[LETTER]++
		}
		if unicode.IsControl(r) {
			count[CONTROL]++
		}
		if unicode.IsGraphic(r) {
			count[GRAPHIC]++
		}
		if unicode.IsDigit(r) {
			count[DIGIT]++
		}
		if unicode.IsMark(r) {
			count[MARK]++
		}
		if unicode.IsSymbol(r) {
			count[SYMBOL]++
		}
		if unicode.IsSpace(r) {
			count[SPACE]++
		}
	}
	fmt.Print("Type\tCount\n")
	fmt.Printf("Letter\t%d\n", count[LETTER])
	fmt.Printf("Control\t%d\n", count[CONTROL])
	fmt.Printf("Graphic\t%d\n", count[GRAPHIC])
	fmt.Printf("Digit\t%d\n", count[DIGIT])
	fmt.Printf("Mark\t%d\n", count[MARK])
	fmt.Printf("Symbol\t%d\n", count[SYMBOL])
	fmt.Printf("Space\t%d\n", count[SPACE])

}
