package types

import (
	"fmt"

	"github.com/go-playground/validator"
)

// PullRequest contains the information about a Pull Request to review
type PullRequest struct {
	RepositoryName    string `json:"repositoryName" validate:"required"`
	RepositoryOwner   string `json:"repositoryOwner" validate:"required"`
	PullRequestNumber int    `json:"pullRequestNumber" validate:"required"`
}

// PullRequestReview contains a review for the Pull Request.
type PullRequestReview struct {
	Review string `json:"pullRequestReview"`
}

// PullRequestValidation is a Valide object for the PullRequest type
type PullRequestValidation struct {
	validate *validator.Validate
}

// PullRequestValidation creates a new Validation for the PullRequest type
func NewPullRequestValidation() *PullRequestValidation {
	validate := validator.New()

	return &PullRequestValidation{validate}
}

// ValidationError wraps the validators FieldError so it is not
// made public
type ValidationError struct {
	validator.FieldError
}

// Error converts a ValidationError to a string
func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// ValidationErrors is a slice of ValidationError
type ValidationErrors []ValidationError

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// Validate validates the given interface and returns any errors with the
// validation
func (v *PullRequestValidation) Validate(i interface{}) ValidationErrors {
	var returnErrs []ValidationError
	if errs, ok := v.validate.Struct(i).(validator.ValidationErrors); ok {
		if errs != nil {
			for _, err := range errs {
				if fe, ok := err.(validator.FieldError); ok {
					ve := ValidationError{fe}
					returnErrs = append(returnErrs, ve)
				}
			}
		}
	}

	return returnErrs
}
