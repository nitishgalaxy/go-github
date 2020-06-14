package services

import (
	"strings"

	"github.com/nitishgalaxy/go-github/src/api/config"
	"github.com/nitishgalaxy/go-github/src/api/providers/github_provider"

	"github.com/nitishgalaxy/go-github/src/api/models/github"
	"github.com/nitishgalaxy/go-github/src/api/models/repositories"
	"github.com/nitishgalaxy/go-github/src/api/utils/errors"
)

type repoService struct{}

type repoServiceInterface interface {
}

var (
	RepoService repoService
)

func init() {
	RepoService = repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewAPiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}

	return &result, nil

}
