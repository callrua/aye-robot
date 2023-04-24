package main

import (
	"aye-robot/pkg/api/services"
	"aye-robot/pkg/api/types"
	"context"
	"log"
	"os"

	"github.com/spf13/pflag"
)

// CLI flags
var (
	repositoryOwner   string
	repositoryName    string
	pullRequestNumber int
)

func main() {
	pflag.StringVarP(&repositoryOwner, "repository-owner", "o", "", "The owner of the repository of which to post the review.")
	pflag.StringVarP(&repositoryName, "repository-name", "n", "", "The name of the repository of which to post the review.")
	pflag.IntVarP(&pullRequestNumber, "pull-request-number", "p", -1, "The ID of the Pull Request of which to review.")
	pflag.Parse()

	ghToken := os.Getenv("GH_TOKEN")
	aiToken := os.Getenv("AI_TOKEN")

	ctx := context.Background()
	ai := services.NewChatGPTClient(ctx, aiToken)
	git := services.NewGithubClientWithAuth(ctx, ghToken)
	logger := log.New(os.Stdout, "aye-robot ", log.LstdFlags)

	// TODO: Add validation for CLI use case, and CLI flags
	// validation := types.NewPullRequestValidation()

	pr := &types.PullRequest{
		RepositoryName:    repositoryName,
		RepositoryOwner:   repositoryOwner,
		PullRequestNumber: pullRequestNumber,
	}

	raw := git.GetRaw(pr.RepositoryOwner, pr.RepositoryName, pr.PullRequestNumber)
	chatGPTReview, err := ai.AskAi(raw)
	if err != nil {
		logger.Println("Error connecting to chatGPT", err)
	}

	prr := &types.PullRequestReview{
		Review: chatGPTReview,
	}
	err = git.PostComment(pr, prr)
	if err != nil {
		logger.Println("Error posting comment to GitHub", err)
	}

}
