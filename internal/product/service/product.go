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

func (ps *ProductService) GetProductByID(ctx context.Context, productID int64) (*entity.ProductWithReview, error) {
	product, err := ps.productRepository.GetByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	reviews, err := ps.productRepository.GetReviewByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}

	return &entity.ProductWithReview{
		Product: *product,
		Reviews: reviews,
	}, nil

}

func (ps *ProductService) ListProducts(
	ctx context.Context,
	listParams entity.ListProductParams,
) ([]entity.Product, error) {
	var createdDescending = new(bool)
	var ratingDescending = new(bool)

	sortCreatedTrimmedLowered := strings.ToLower(strings.TrimSpace(listParams.SortCreated))
	sortRatingTrimmedLowered := strings.ToLower(strings.TrimSpace(listParams.SortRating))

	if sortCreatedTrimmedLowered == "oldest" {
		*createdDescending = false
	} else if sortCreatedTrimmedLowered == "newest" {
		*createdDescending = true
	}

	if sortRatingTrimmedLowered == "lowest" {
		*ratingDescending = false
	} else if sortRatingTrimmedLowered == "highest" {
		*ratingDescending = true
	}

	sku := strings.TrimSpace(listParams.Sku)
	title := strings.TrimSpace(listParams.Title)
	category := strings.TrimSpace(listParams.Category)
	etalase := strings.TrimSpace(listParams.Etalase)
	filter := entity.FilterList{
		Sku:               sku,
		Title:             title,
		Category:          category,
		Etalase:           etalase,
		CreatedDescending: createdDescending,
		RatingDescending:  ratingDescending,
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
	product, err := ps.productRepository.GetByID(ctx, params.ID)
	if err != nil {
		return err
	}

	updatedProduct := params.UpdateProduct(product)

	return ps.productRepository.Update(ctx, updatedProduct)
}
