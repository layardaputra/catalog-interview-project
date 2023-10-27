package app

import (
	"github.com/jmoiron/sqlx"
	product_v1 "github.com/layardaputra/govtech-catalog-test-project/internal/product/api/v1/handlers"
	productRepository "github.com/layardaputra/govtech-catalog-test-project/internal/product/domain/repository"
	productService "github.com/layardaputra/govtech-catalog-test-project/internal/product/service"
)

type App struct {
	ProductHandler *product_v1.HandlerV1
}

func NewApp(db *sqlx.DB) *App {
	// Initialize service
	productRepo := productRepository.NewSqlRepository(db)
	productServ := productService.NewProductService(productRepo)
	productHandler := product_v1.NewHandlerV1(db, productServ)

	return &App{
		ProductHandler: productHandler,
	}
}
