package main

import "fmt"

// re-write reverse to use an array pointer instead of a slice

func reverse(s *[5]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	var a = [...]int{1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Printf("%v", a)
}
