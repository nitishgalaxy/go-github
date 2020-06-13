package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nitishgalaxy/go-github/src/api/clients/restclient"
	"github.com/nitishgalaxy/go-github/src/api/models/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(access_token string) string {
	return fmt.Sprintf(headerAuthorizationFormat, access_token)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))
	response, err := restclient.Post(urlCreateRepo, request, headers)
	fmt.Println("Response from github :", response)

	if err != nil {
		log.Printf("Error creating new repo in github: %s", err.Error())
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "Invalid response body."}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "Invalid json response body."}
		}
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Printf("Error trying to unmarshal github Create Repo successful Response: %s", err.Error())
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "Error trying to unmarshal github Create Repo successful Response."}
	}

	return &github.CreateRepoResponse{}, nil
}
