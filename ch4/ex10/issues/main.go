package main

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/artgillespie/gopl/ch4/github"
)

func String(i *github.Issue) string {
	return fmt.Sprintf("#%-5d %-55.55s %-9.9s %-8s %s\n", i.Number, i.Title, i.User.Login, i.State, i.CreatedAt)
}

func PrintCategory(header string, issues github.IssueSlice) {
	fmt.Println("-------------------------")
	fmt.Printf("%s: %d\n", header, len(issues))
	fmt.Println("-------------------------")
	for _, issue := range issues {
		fmt.Print(String(issue))
	}
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		fmt.Errorf("Couldn't search GitHub Issues: %s", err)
		os.Exit(-1)
	}
	sort.Sort(sort.Reverse(result.Items))
	var oneMonth, oneYear, moreThanOneYear github.IssueSlice
	oneMonthAgo := time.Now().AddDate(0, -1, 0)
	oneYearAgo := time.Now().AddDate(-1, 0, 0)
	fmt.Printf("num items %d\n", len(result.Items))
	for _, issue := range result.Items {
		var category = &moreThanOneYear
		if issue.CreatedAt.After(oneMonthAgo) {
			category = &oneMonth
		} else if issue.CreatedAt.After(oneYearAgo) {
			category = &oneYear
		}
		*category = append(*category, issue)
	}
	PrintCategory("Issues in the past month", oneMonth)
	PrintCategory("Issues in the past year", oneYear)
	PrintCategory("Issues older than a year", moreThanOneYear)
}
