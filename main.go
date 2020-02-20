// Package main implements a word counting CLI application
package main

import (
	"flag"
	"fmt"
	"wordcounter/fetchapi"
)

// main is the entry point for the application
func main() {
	fmt.Println("Starting Word Counter")

	var commentId string
	var postId string

	flag.StringVar(&commentId, "commentId", "0", "the comment ID")
	flag.StringVar(&postId, "postId", "0", "the post ID")
	flag.Parse()

	fmt.Println(commentId, postId)

	api := &fetchapi.FetchApi{
		PostId:    postId,
		CommentId: commentId,
	}

	comments := api.GetComments()

	fmt.Println(comments)


}
