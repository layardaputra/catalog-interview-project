package repository

import (
	"context"

	"github.com/layardaputra/govtech-catalog-test-project/internal/product/domain/entity"
)

// IProductRepository defines the contract for managing product data.
type IProductRepository interface {
	// product related
	GetByID(ctx context.Context, productID int64) (*entity.Product, error)
	List(ctx context.Context, filter entity.FilterList) ([]entity.Product, error)
	Create(ctx context.Context, product *entity.Product) error
	Update(ctx context.Context, product *entity.Product) error

	// product review related
	CreateReview(ctx context.Context, review *entity.Review) error
	GetReviewByProductID(ctx context.Context, productID int64) ([]entity.Review, error)
}
