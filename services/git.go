package services

import "aye-robot/types"

type GitClient interface {
	GetRaw(string, string, int) *string
	PostComment(*types.PullRequest, *types.PullRequestReview) error
}
