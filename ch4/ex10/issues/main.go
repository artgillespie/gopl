package main

import (
	"fmt"
	"os"

	"github.com/artgillespie/gopl/ch4/ex10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		fmt.Errorf("Couldn't search GitHub Issues: %s", err)
		os.Exit(-1)
	}
	for _, issue := range result.Items {
		fmt.Printf("%d\t%s\t%s\t%s\n", issue.Number, issue.Title, issue.User.Login, issue.State)
	}
}
