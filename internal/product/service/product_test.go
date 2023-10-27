package service_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/layardaputra/govtech-catalog-test-project/common"
	"github.com/layardaputra/govtech-catalog-test-project/internal/product/domain/entity"
	"github.com/layardaputra/govtech-catalog-test-project/internal/product/service"
	"github.com/stretchr/testify/assert"
)

type mockProductRepository struct {
}

func (m *mockProductRepository) GetByID(ctx context.Context, productID int64) (*entity.Product, error) {
	var id int64 = 2
	if productID == 2 {
		return &entity.Product{
			ID:          &id,
			Sku:         "test sku",
			Title:       "test title",
			Description: "test description",
			Category:    "test category",
			Etalase:     "test etalase",
			Images: []entity.ProductImage{
				{
					ImageURL:    "test_image_1",
					Description: "test_desc_1",
				},
			},
			Weight:      1.2,
			Price:       10000,
			ReviewCount: 1,
			TotalRating: 5,
			AVGRating:   5,
			CreatedAt:   time.Now(),
		}, nil
	}

	return nil, &common.CustomError{
		StatusCode: http.StatusNotFound,
		Message:    "Data Not Found",
		Err:        nil,
	}
}

func (m *mockProductRepository) List(ctx context.Context, filter entity.FilterList) ([]entity.Product, error) {
	if filter.Sku == "no data" {
		return []entity.Product{}, nil
	}

	var id int64 = 1
	return []entity.Product{
		{
			ID:          &id,
			Sku:         "test sku",
			Title:       "test title",
			Description: "test description",
			Category:    "test category",
			Etalase:     "test etalase",
			Images: []entity.ProductImage{
				{
					ImageURL:    "test_image_1",
					Description: "test_desc_1",
				},
			},
			Weight:      1.2,
			Price:       10000,
			ReviewCount: 1,
			TotalRating: 5,
			AVGRating:   5,
			CreatedAt:   time.Now(),
		},
	}, nil
}

func (m *mockProductRepository) Create(ctx context.Context, product *entity.Product) error {
	return nil
}

func (m *mockProductRepository) Update(ctx context.Context, product *entity.Product) error {
	return nil
}

func (m *mockProductRepository) CreateReview(ctx context.Context, review *entity.Review) error {
	return nil
}

func (m *mockProductRepository) GetReviewByProductID(ctx context.Context, productID int64) ([]entity.Review, error) {
	if productID == 1 {
		return nil, &common.CustomError{
			StatusCode: http.StatusNotFound,
			Message:    "Data Not Found",
			Err:        nil,
		}
	}

	var id int64 = 2
	return []entity.Review{
		{
			ID:            &id,
			ProductID:     2,
			Rating:        5,
			ReviewComment: "mantap",
			CreatedAt:     time.Now(),
		},
	}, nil
}

func TestGetProductByID(t *testing.T) {
	mockRepo := &mockProductRepository{}

	ps := service.NewProductService(mockRepo)

	product, err := ps.GetProductByID(context.Background(), 2)

	assert.NoError(t, err)
	assert.NotNil(t, product)
}

func TestGetProductByIDNotFound(t *testing.T) {
	mockRepo := &mockProductRepository{}

	ps := service.NewProductService(mockRepo)

	product, err := ps.GetProductByID(context.Background(), 1)

	assert.Error(t, err)
	assert.Nil(t, product)
}

func TestListProducts(t *testing.T) {
	mockRepo := &mockProductRepository{}

	ps := service.NewProductService(mockRepo)

	product, err := ps.ListProducts(context.Background(), entity.ListProductParams{
		Title: "some random title",
	})

	assert.NoError(t, err)
	assert.NotNil(t, product)
}

func TestListProductsEmptyList(t *testing.T) {
	mockRepo := &mockProductRepository{}

	ps := service.NewProductService(mockRepo)

	product, err := ps.ListProducts(context.Background(), entity.ListProductParams{
		Sku: "no data",
	})

	assert.NoError(t, err)
	assert.Empty(t, product)
}

func TestCreateProduct(t *testing.T) {
	mockRepo := &mockProductRepository{}

	ps := service.NewProductService(mockRepo)

	err := ps.CreateProduct(context.Background(), entity.CreateProductParams{
		Sku:         "test_new_sku",
		Title:       "test_new_title",
		Description: "test_description",
		Category:    "test_category",
		Etalase:     "test_etalase",
		Images:      []entity.ProductImage{},
		Weight:      1,
		Price:       10000,
	})

	assert.NoError(t, err)
}

func TestUpdateProduct(t *testing.T) {
	mockRepo := &mockProductRepository{}

	ps := service.NewProductService(mockRepo)
	sku := "test_update_sku"

	err := ps.UpdateProduct(context.Background(), entity.UpdateProductParams{
		ID:  2,
		Sku: &sku,
	})

	assert.NoError(t, err)
}

func TestUpdateProductNotFoundProduct(t *testing.T) {
	mockRepo := &mockProductRepository{}

	ps := service.NewProductService(mockRepo)
	sku := "test_update_sku"

	err := ps.UpdateProduct(context.Background(), entity.UpdateProductParams{
		ID:  1,
		Sku: &sku,
	})

	assert.Error(t, err)
}
