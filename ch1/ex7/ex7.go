package main

import (
	"fmt"
	"io"
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
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Printf("Error copying response body to stdout for url %s: %v", url, err)
		}
	}
}
