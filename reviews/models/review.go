package models

import (
	"errors"
	"time"
)

const maxLengthComment = 400

// Review represent an anon review from some website
type Review struct {
	ID      int64
	Stars   int8
	Comment string
	Date    time.Time
}

// CreateReviewCMD command to create a new review
type CreateReviewCMD struct {
	Stars   int8   `json:"stars"`
	Comment string `json:"comment"`
}

func (cmd *CreateReviewCMD) validate() error {
	if cmd.Stars < 1 || cmd.Stars > 5 {
		return errors.New("stars must be between 0 and 5")
	}

	if len(cmd.Comment) > maxLengthComment {
		return errors.New("comment must be less than 400 characters")
	}

	return nil
}
