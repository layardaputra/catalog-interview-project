package requests

import (
	"net/http"

	"github.com/layardaputra/govtech-catalog-test-project/common"
)

type AddProductReviewRequest struct {
	ProductID     int64
	Rating        uint8  `json:"rating"`
	ReviewComment string `json:"review_comment"`
}

func (req *AddProductReviewRequest) ValidateRequest() error {
	if req.Rating < 1 && req.Rating > 5 {
		return &common.CustomError{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Invalid Rating, Rating range between 1 - 5",
			Err:        nil,
		}
	}

	if len(req.ReviewComment) > 200 {
		return &common.CustomError{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Maximum Letter of Review Comment is 200 Letters",
			Err:        nil,
		}
	}

	return nil
}
