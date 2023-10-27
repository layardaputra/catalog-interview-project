package repository

import (
	"context"

	"github.com/layardaputra/govtech-catalog-test-project/internal/product/domain/entity"
)

// IProductRepository defines the contract for managing product data.
type IProductRepository interface {
	GetByID(ctx context.Context, productID int64) (*entity.Product, error)
	List(ctx context.Context, filter entity.FilterList) ([]entity.Product, error)
	Create(ctx context.Context, product *entity.Product) error
	Update(ctx context.Context, product *entity.Product) error
}
