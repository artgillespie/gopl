package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error fetching URL %s: %v\n", url, err)
			os.Exit(-1)
		}
		defer resp.Body.Close()
		fmt.Printf("Fetched URL %s with status %s\n", url, resp.Status)
	}
}
