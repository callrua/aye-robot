package services

import (
	"aye-robot/types"
	"context"
	"log"

	"github.com/google/go-github/v51/github"
	"golang.org/x/oauth2"
)

type GithubClient struct {
	client *github.Client
}

// NewGithubClientWithAuth returns a GitHub client with oauth2 authentication
func NewGithubClientWithAuth(ctx context.Context, token string) *GithubClient {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	oauth := oauth2.NewClient(ctx, ts)

	return &GithubClient{
		client: github.NewClient(oauth),
	}
}

// GetRaw returns the raw diff of a Pull Request
func (g *GithubClient) GetRaw(owner string, repository string, pullRequestNumber int) *string {
	rawOpts := github.RawOptions{
		Type: github.Diff,
	}
	raw, resp, err := g.client.PullRequests.GetRaw(context.TODO(), owner, repository, pullRequestNumber, rawOpts)
	if err != nil {
		log.Printf("Unable to Get Raw PR, got response %v, and error %v", resp, err)
	}

	return &raw
}

// PostComment posts a comment to a GitHub Pull Request
func (g *GithubClient) PostComment(pr *types.PullRequest, prr *types.PullRequestReview) error {
	// Post a comment to the PR
	comment := github.IssueComment{
		Body: &prr.Review,
	}

	_, _, err := g.client.Issues.CreateComment(context.TODO(), pr.RepositoryOwner, pr.RepositoryName, pr.PullRequestNumber, &comment)
	if err != nil {
		return err
	}

	return nil
}
