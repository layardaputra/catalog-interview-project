package service

import (
	"context"
	"strings"

	"github.com/layardaputra/govtech-catalog-test-project/internal/product/domain/entity"
	"github.com/layardaputra/govtech-catalog-test-project/internal/product/domain/repository"
)

type ProductService struct {
	productRepository repository.IProductRepository
}

func NewProductService(repo repository.IProductRepository) *ProductService {
	return &ProductService{
		productRepository: repo,
	}
}

func (ps *ProductService) GetProductByID(ctx context.Context, productID int64) (*entity.Product, error) {
	return ps.productRepository.GetByID(ctx, productID)
}

func (ps *ProductService) ListProducts(
	ctx context.Context,
	listParams entity.ListProductParams,
) ([]entity.Product, error) {
	isDescending := true
	sortTrimmedLowered := strings.ToLower(strings.TrimSpace(listParams.Sort))

	if sortTrimmedLowered == "oldest" {
		isDescending = false
	}

	sku := strings.TrimSpace(listParams.Sku)
	title := strings.TrimSpace(listParams.Title)
	category := strings.TrimSpace(listParams.Category)
	etalase := strings.TrimSpace(listParams.Etalase)
	filter := entity.FilterList{
		Sku:          sku,
		Title:        title,
		Category:     category,
		Etalase:      etalase,
		IsDescending: isDescending,
	}

	return ps.productRepository.List(ctx, filter)
}

func (ps *ProductService) CreateProduct(ctx context.Context, params entity.CreateProductParams) error {
	product := entity.NewProduct(
		params.Sku,
		params.Title,
		params.Description,
		params.Category,
		params.Etalase,
		params.Images,
		params.Weight,
		params.Price,
	)

	return ps.productRepository.Create(ctx, &product)
}

func (ps *ProductService) UpdateProduct(ctx context.Context, params entity.UpdateProductParams) error {
	product, err := ps.GetProductByID(ctx, params.ID)
	if err != nil {
		return err
	}

	updatedProduct := params.UpdateProduct(product)

	return ps.productRepository.Update(ctx, updatedProduct)
}
