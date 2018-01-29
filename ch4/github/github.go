package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count`
	Items      IssueSlice
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type IssueSlice []*Issue

func (issues IssueSlice) Len() int {
	return len(issues)
}

func (issues IssueSlice) Less(i, j int) bool {
	return issues[i].CreatedAt.Before(issues[j].CreatedAt)
}

func (issues IssueSlice) Swap(i, j int) {
	issues[i], issues[j] = issues[j], issues[i]
}

// SearchIssues queries the GitHub issue tracker

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error getting issues %d %s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result IssueSearchResult
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func CreateIssue(title string, body string) (*Issue, error) {
	var p = struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}{title, body}
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	res, err := http.Post(IssuesURL+"/repos/artgillespie/gopl/issues", "application/json", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Issue not created %d %s", res.StatusCode, res.Status)
	}
	defer res.Body.Close()
	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var issue Issue
	err = json.Unmarshal(b, &issue)
	if err != nil {
		return nil, err
	}
	return &issue, nil
}
