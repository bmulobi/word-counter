// Package main implements a word counting CLI application
package main

import (
	"flag"
	"fmt"
	"wordcounter/counter"
	"wordcounter/fetchapi"
)

// main is the entry point for the application
func main() {
	fmt.Println("\nStarting Word Counter")

	var commentId string
	var postId string

	flag.StringVar(&commentId, "commentId", "0", "the comment ID")
	flag.StringVar(&postId, "postId", "0", "the post ID")
	flag.Parse()
	api := &fetchapi.Arguments{
		PostId:    postId,
		CommentId: commentId,
	}

	comments := api.GetComments()
	count := counter.Counter{WordCounts: map[string]int{}}
	count.CountWords(comments)

	count.DisplayLeastUsed()
}
