package types

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToJSON(t *testing.T) {
	prr := PullRequestReview{
		Review: "This is a review",
	}
	b := bytes.NewBufferString("")
	err := ToJSON(prr, b)
	assert.Empty(t, err)
}

func TestFromJSON(t *testing.T) {
	pr := &PullRequest{}
	in := map[string]interface{}{"repositoryOwner": "owner", "repositoryName": "my-repo", "pullRequestNumber": 1}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(in)

	err := FromJSON(pr, b)

	assert.Empty(t, err)
	assert.Equal(t, "owner", pr.RepositoryOwner)
	assert.Equal(t, "my-repo", pr.RepositoryName)
	assert.Equal(t, 1, pr.PullRequestNumber)
}
