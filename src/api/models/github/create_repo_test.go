package github

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "nitish-go-test-repo",
		Description: "A golang test repo",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	// Marshal takes an input interface and attempts to create a valid JSON string.
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, `{"name":"nitish-go-test-repo","description":"A golang test repo","homepage":"","private":true,"has_issues":true,"has_projects":true,"has_wiki":true}`, string(bytes))
	//assert.EqualValues(t, )

	fmt.Println(string(bytes))

	// Convert the previous json and load it in 'target' struct variable.
	var target CreateRepoRequest
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)

	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)

}
