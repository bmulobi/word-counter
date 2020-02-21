// Package fetchapi implements methods for fetching comments from
// https://jsonplaceholder.typicode.com/comments/
package fetchapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Arguments contains the query parameters (if any) for fetching comments
type Arguments struct {
	PostId    string
	CommentId string
}

// Comment holds only the comment string we are interested in
// this will ensure we discard the other fields from the API response
// during decoding
type Comment struct {
	Body string
}

// Get comments from the API given the url
func (api *Arguments) GetComments() []Comment {
	url, err := api.buildUrl()
	if err != nil {
		fmt.Println("Error trying to build url : ", err)
		log.Fatal(err)
	}

	client := http.Client{Timeout: 10 * time.Second}
	response, err := client.Get(url)
	if err != nil {
		fmt.Println("Error trying to fetch comments : ", err)
		log.Fatal(err)
	}
	defer response.Body.Close()

	var comments []Comment
	json.NewDecoder(response.Body).Decode(&comments)

	return comments
}

func (api *Arguments) buildUrl() (string, error) {
	var commentIdInt int
	var postIdInt int

	_, err := fmt.Sscanf(api.CommentId, "%d", &commentIdInt)
	_, err = fmt.Sscanf(api.PostId, "%d", &postIdInt)

	if err != nil {
		return "Conversion failure", err
	}

	var url string
	switch {
	case postIdInt != 0 && commentIdInt != 0:
		url = fmt.Sprintf(
			"https://jsonplaceholder.typicode.com/comments/?postId=%s&id=%s",
			api.PostId, api.CommentId,
		)
	case postIdInt != 0:
		url = fmt.Sprintf("https://jsonplaceholder.typicode.com/comments/?postId=%s", api.PostId)
	case commentIdInt != 0:
		url = fmt.Sprintf("https://jsonplaceholder.typicode.com/comments/?id=%s", api.CommentId)
	default:
		url = "https://jsonplaceholder.typicode.com/comments/"
	}

	return url, nil
}
