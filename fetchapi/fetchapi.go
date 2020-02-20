// Package fetchapi implements methods for fetching comments from
// https://jsonplaceholder.typicode.com/comments/
package fetchapi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// FetchApi contains the query parameters (if any) for fetching comments
type FetchApi struct {
	PostId string
	CommentId string
}

// Get comments from the API given the url
func (api *FetchApi) GetComments() string {

	url, err := api.buildUrl()

	if err != nil {
		fmt.Println("Error trying to build url : ", err)
		log.Fatal(err)
	}

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error trying to fetch comments : ", err)
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	return string(body)
}

func (api *FetchApi) buildUrl() (string, error) {
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
			api.PostId, api.CommentId)
	case postIdInt != 0:
		url = fmt.Sprintf("https://jsonplaceholder.typicode.com/comments/?postId=%s", api.PostId)
	case commentIdInt != 0:
		url = fmt.Sprintf("https://jsonplaceholder.typicode.com/comments/?id=%s", api.CommentId)
	default:
		url = "https://jsonplaceholder.typicode.com/comments/"
	}

	return url, nil
}
