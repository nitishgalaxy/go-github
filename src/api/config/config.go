package config

import "os"

const (
	apiGitHibAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv(apiGitHibAccessToken)
)

func GetGithubAccessToken() string {
	return githubAccessToken
}
