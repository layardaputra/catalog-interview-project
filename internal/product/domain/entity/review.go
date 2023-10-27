package entity

import "time"

type Review struct {
	ID            *int64    `json:"id" db:"id"`
	ProductID     int64     `json:"product_id" db:"product_id"`
	Rating        uint8     `json:"rating" db:"rating"`
	ReviewComment string    `json:"review_comment" db:"review_comment"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

type CreateReviewParam struct {
	ProductID     int64
	Rating        uint8
	ReviewComment string
}
