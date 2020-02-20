// Package counter implements a counter for the least used words
package counter

import "fmt"

// Counter contains a map of all the word counts
type Counter struct {
	WordCounts map[string]int
}

// CountWords gets the count for each word in the comments
func (counter *Counter) CountWords() map[string]int {

	return map[string]int{}
}

// DisplayLeastUsed gets the four least used words and displays them with their counts
func (counter *Counter) DisplayLeastUsed() {
	for word, count := range counter.WordCounts {
		fmt.Println(word, count)
	}
}