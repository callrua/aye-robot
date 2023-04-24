package handlers

import (
	"aye-robot/pkg/api/types"
	"context"
	"net/http"
)

func (p *PullRequestReviewHandler) MiddlewareValidatePullRequestReview(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		pr := &types.PullRequest{}

		err := types.FromJSON(pr, r.Body)
		if err != nil {
			p.logger.Println("[ERROR] deserializing pr", err)

			rw.WriteHeader(http.StatusBadRequest)
			types.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		// validate the pr
		errs := p.validation.Validate(pr)
		if len(errs) != 0 {
			p.logger.Println("[ERROR] validating pr", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			types.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		// add the pr to the context
		ctx := context.WithValue(r.Context(), KeyPullRequest{}, pr)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
