package models

import (
	"github.com/jaswdr/faker"
	"testing"
)

func NewReview(start int8, comment string) *CreateReviewCMD {
	return &CreateReviewCMD{
		Stars:   start,
		Comment: comment,
	}
}

func Test_Create_Review_Validate_Assert_True(t *testing.T) {
	r := NewReview(5, "The new iphone X looks good")
	if err := r.validate(); err != nil {
		t.Error("The validation did not pass ", err)
		t.Fail()
	}
}

func Test_Create_Review_Validate_Assert_Fail_For_Number_Start(t *testing.T) {
	r := NewReview(8, "The new iphone X looks good")
	if err := r.validate(); err == nil {
		t.Error("should fail with 5 stars")
		t.Fail()
	}
}

func Test_Create_Review_Validate_Assert_Fail_For_Long_Message(t *testing.T) {
	r := NewReview(4, faker.New().Lorem().Text(500))
	if err := r.validate(); err == nil {
		t.Error("should fail with long message")
		t.Fail()
	}
}
