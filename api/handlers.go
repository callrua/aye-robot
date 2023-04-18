package api

import (
	"aye-robot/services"
	"aye-robot/types"
	"log"
	"net/http"
)

// KeyPullRequest is a key used for the PullRequest object in the context of a request
type KeyPullRequest struct{}

// GenericError is a generic error message returned by the server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

type PullRequestReviewHandler struct {
	aiClient   services.AiClient
	gitClient  services.GitClient
	logger     *log.Logger
	validation *types.PullRequestValidation
}

func NewPullRequestReviewHandler(a services.AiClient, g services.GitClient, l *log.Logger, v *types.PullRequestValidation) *PullRequestReviewHandler {
	return &PullRequestReviewHandler{
		aiClient:   a,
		gitClient:  g,
		logger:     l,
		validation: v,
	}
}

func (p *PullRequestReviewHandler) ReviewPR(rw http.ResponseWriter, r *http.Request) {
	pr := r.Context().Value(KeyPullRequest{}).(*types.PullRequest)

	p.logger.Printf("Received request for pull request %d, for repository %s, under owner %s\n", pr.PullRequestNumber, pr.RepositoryName, pr.RepositoryOwner)

	raw := p.gitClient.GetRaw(pr.RepositoryOwner, pr.RepositoryName, pr.PullRequestNumber)
	// TODO: make this debug logging
	// log.Printf(*raw)

	chatGPTReview, err := p.aiClient.AskAi(raw)
	if err != nil {
		http.Error(rw, "Error connecting to chatGPT", http.StatusInternalServerError)
	}

	prr := &types.PullRequestReview{
		Review: chatGPTReview,
	}

	err = types.ToJSON(prr, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}

	p.gitClient.PostComment(pr, prr)
}
