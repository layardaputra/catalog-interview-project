package product_v1

import (
	"github.com/jmoiron/sqlx"
	"github.com/layardaputra/govtech-catalog-test-project/internal/product/service"
)

type HandlerV1 struct {
	DB             *sqlx.DB
	ProductService *service.ProductService
}

func NewHandlerV1(db *sqlx.DB, svc *service.ProductService) *HandlerV1 {
	return &HandlerV1{
		DB:             db,
		ProductService: svc,
	}
}
