package entity

import (
	"time"

	"github.com/jmoiron/sqlx/types"
)

type Product struct {
	ID          *int64         `json:"id"`
	Sku         string         `json:"sku"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Etalase     string         `json:"etalase"`
	Images      []ProductImage `json:"images"`
	Weight      float64        `json:"weight"`
	Price       float64        `json:"price"`
	ReviewCount uint64         `json:"-"`
	TotalRating uint64         `json:"-"`
	AVGRating   float64        `json:"rating"`
	CreatedAt   time.Time      `json:"created_at"`
}

type ProductWithReview struct {
	Product
	Reviews []Review `json:"reviews"`
}

type ProductSQLX struct {
	ID          *int64         `db:"id"`
	Sku         string         `db:"sku"`
	Title       string         `db:"title"`
	Description string         `db:"description"`
	Category    string         `db:"category"`
	Etalase     string         `db:"etalase"`
	Images      types.JSONText `db:"images"`
	Weight      float64        `db:"weight"`
	Price       float64        `db:"price"`
	ReviewCount uint64         `db:"review_count"`
	TotalRating uint64         `db:"total_rating"`
	CreatedAt   time.Time      ` db:"created_at"`
}

type ProductImage struct {
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}

func NewProduct(
	sku string,
	title string,
	description string,
	category string,
	etalase string,
	images []ProductImage,
	weight float64,
	price float64,
) Product {
	return Product{
		ID:          nil,
		Sku:         sku,
		Title:       title,
		Description: description,
		Category:    category,
		Etalase:     etalase,
		Images:      images,
		Weight:      weight,
		Price:       price,
		AVGRating:   0,
	}
}

type FilterList struct {
	Sku               string
	Title             string
	Category          string
	Etalase           string
	CreatedDescending *bool
	RatingDescending  *bool
}

type ListProductParams struct {
	Sku         string
	Title       string
	Category    string
	Etalase     string
	SortCreated string
	SortRating  string
}

type CreateProductParams struct {
	Sku         string         `json:"sku"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Etalase     string         `json:"etalase"`
	Images      []ProductImage `json:"images"`
	Weight      float64        `json:"weight"`
	Price       float64        `json:"price"`
}

type UpdateProductParams struct {
	ID          int64
	Sku         *string        `json:"sku"`
	Title       *string        `json:"title"`
	Description *string        `json:"description"`
	Category    *string        `json:"category"`
	Etalase     *string        `json:"etalase"`
	Images      []ProductImage `json:"images"`
	Weight      *float64       `json:"weight"`
	Price       *float64       `json:"price"`
}

func (upp UpdateProductParams) UpdateProduct(product *Product) *Product {
	if product.ID == nil {
		return product
	}

	if *product.ID != upp.ID {
		return product
	}

	if upp.Sku != nil && *upp.Sku != "" {
		product.Sku = *upp.Sku
	}

	if upp.Title != nil && *upp.Title != "" {
		product.Title = *upp.Title
	}

	if upp.Description != nil && *upp.Description != "" {
		product.Description = *upp.Description
	}

	if upp.Category != nil && *upp.Category != "" {
		product.Category = *upp.Category
	}

	if upp.Etalase != nil && *upp.Etalase != "" {
		product.Etalase = *upp.Etalase
	}

	if upp.Weight != nil && *upp.Weight > 0 {
		product.Weight = *upp.Weight
	}

	if upp.Price != nil && *upp.Weight > 0 {
		product.Price = *upp.Price
	}

	if len(upp.Images) != 0 {
		product.Images = append(product.Images, upp.Images...)
	}

	return product
}
