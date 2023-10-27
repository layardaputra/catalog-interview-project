package service_test

import (
	"context"
	"testing"

	"github.com/layardaputra/govtech-catalog-test-project/internal/product/domain/entity"
	"github.com/layardaputra/govtech-catalog-test-project/internal/product/service"
	"github.com/stretchr/testify/assert"
)

func TestAddReviewProduct(t *testing.T) {
	mockRepo := &mockProductRepository{}

	ps := service.NewProductService(mockRepo)
	err := ps.AddReviewProduct(context.Background(), entity.CreateReviewParam{
		ProductID:     2,
		Rating:        1,
		ReviewComment: "mantapp",
	})

	assert.NoError(t, err)
}
