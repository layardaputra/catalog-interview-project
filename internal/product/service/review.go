package service

import (
	"context"

	"github.com/layardaputra/govtech-catalog-test-project/internal/product/domain/entity"
)

func (ps *ProductService) AddReviewProduct(ctx context.Context, param entity.CreateReviewParam) error {
	product, err := ps.productRepository.GetByID(ctx, param.ProductID)
	if err != nil {
		return err
	}

	newReview := entity.Review{
		ProductID:     param.ProductID,
		Rating:        param.Rating,
		ReviewComment: param.ReviewComment,
	}

	err = ps.productRepository.CreateReview(
		ctx,
		&newReview,
	)
	if err != nil {
		return err
	}

	product.ReviewCount += 1
	product.TotalRating += uint64(newReview.Rating)

	err = ps.productRepository.Update(ctx, product)
	if err != nil {
		return err
	}

	return nil
}
