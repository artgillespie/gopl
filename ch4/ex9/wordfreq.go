package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// sort by frequency
type WordFreq struct {
	Word  string
	Count int
}

type WordFreqSlice []WordFreq

func (a WordFreqSlice) Len() int           { return len(a) }
func (a WordFreqSlice) Less(i, j int) bool { return a[i].Count > a[j].Count }
func (a WordFreqSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

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

	// sort in descending order by number of instances
	var wordFreqs WordFreqSlice
	for k, v := range count {
		wordFreqs = append(wordFreqs, WordFreq{k, v})
	}
	sort.Sort(wordFreqs)
	for _, f := range wordFreqs {
		fmt.Printf("%s\t%d\n", f.Word, f.Count)
	}
}
