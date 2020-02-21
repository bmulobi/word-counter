// Package counter implements a counter for the least used words
package counter

import (
	"fmt"
	"sort"
	"strings"
	"wordcounter/fetchapi"
)

// Counter contains a map of all the word counts
type Counter struct {
	WordCounts map[string]int
}

// CountWords gets the count for each word in the comments
func (counter *Counter) CountWords(comments []fetchapi.Comment) {
	for _, comment := range comments {
		words := strings.Fields(comment.Body)

		for _, word := range words {
			counter.WordCounts[word]++
		}
	}
}

// getLeastUsedWords returns the 4 least used words in the comment(s)
// there maybe more than 4 words that are least used, e.g 10 words with count 1,
// this function will return an arbitrary 4 of those words that are least used
func (counter *Counter) getLeastUsedWords() map[string]int {
	leastUsed := make(map[string]int)
	wordCounts := make([]int, 0, len(counter.WordCounts))

	for _, count := range counter.WordCounts {
		wordCounts = append(wordCounts, count)
	}

	sort.Ints(wordCounts)
	var wordExists bool
	for index := 0; index < 4; index++ {
		for word, count := range counter.WordCounts {
			_, wordExists = leastUsed[word]
			if count == wordCounts[index] && !wordExists {
				leastUsed[word] = count
				break
			}
		}
	}

	return leastUsed
}

// DisplayLeastUsed gets the four least used words and displays them with their counts
func (counter *Counter) DisplayLeastUsed() {
	leastUsedWords := counter.getLeastUsedWords()

	for word, count := range leastUsedWords {
		fmt.Println(word, count)
	}
	fmt.Println()
}
