package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPullRequestType(t *testing.T) {
	pr := PullRequest{
		RepositoryName:    "my-repo",
		RepositoryOwner:   "owner",
		PullRequestNumber: 1,
	}

	v := NewPullRequestValidation()
	err := v.Validate(pr)
	assert.Empty(t, err)
}

func TestEmptyPullRequestReturnsErr(t *testing.T) {
	pr := PullRequest{}

	v := NewPullRequestValidation()
	err := v.Validate(pr)
	assert.Len(t, err, 3)
}

func TestUndefinedRepositoryNameReturnsErr(t *testing.T) {
	pr := PullRequest{
		RepositoryOwner:   "owner",
		PullRequestNumber: 1,
	}

	v := NewPullRequestValidation()
	err := v.Validate(pr)
	assert.Len(t, err, 1)
}

func TestUndefinedRepositoryOwnerReturnsErr(t *testing.T) {
	pr := PullRequest{
		RepositoryName:    "my-repo",
		PullRequestNumber: 1,
	}

	v := NewPullRequestValidation()
	err := v.Validate(pr)
	assert.Len(t, err, 1)
}

func TestUndefinedPullRequestNumberReturnsErr(t *testing.T) {
	pr := PullRequest{
		RepositoryName:  "my-repo",
		RepositoryOwner: "owner",
	}

	v := NewPullRequestValidation()
	err := v.Validate(pr)
	assert.Len(t, err, 1)
}
