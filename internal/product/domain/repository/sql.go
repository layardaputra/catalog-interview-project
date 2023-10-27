package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/layardaputra/govtech-catalog-test-project/common"
	"github.com/layardaputra/govtech-catalog-test-project/internal/product/domain/entity"
)

type SqlProductRepository struct {
	db *sqlx.DB
}

func NewSqlRepository(db *sqlx.DB) IProductRepository {
	return &SqlProductRepository{
		db: db,
	}
}

func (r *SqlProductRepository) GetByID(ctx context.Context, productID int64) (*entity.Product, error) {
	var models []entity.Product

	tx, err := common.GetTransactionFromContext(ctx, r.db)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT
			pr.id,
			pr.sku,
			pr.title,
			pr.description,
			pr.category,
			pr.etalase,
			pr.images,
			pr.weight,
			pr.price,
			pr.review_count,
			pr.total_rating,
			pr.created_at
		FROM product pr
		WHERE pr.id = :id
		ORDER BY 
			pr.created_at DESC
		LIMIT 1;
	`
	params := map[string]interface{}{"id": productID}

	rows, err := tx.NamedQuery(
		query,
		params,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var result entity.ProductSQLX
		if err := rows.StructScan(&result); err != nil {
			return nil, err
		}
		model, err := r.modelToEntity(result)
		if err != nil {
			return nil, err
		}

		models = append(models, *model)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(models) < 1 {
		return nil, &common.CustomError{
			StatusCode: http.StatusNotFound,
			Message:    "Product Not Found",
			Err:        nil,
		}
	}

	return &models[0], nil
}

func (r *SqlProductRepository) List(ctx context.Context, filter entity.FilterList) ([]entity.Product, error) {
	tx, err := common.GetTransactionFromContext(ctx, r.db)
	if err != nil {
		return nil, err
	}

	var models = []entity.Product{}

	query := `
		SELECT
			pr.id,
			pr.sku,
			pr.title,
			pr.description,
			pr.category,
			pr.etalase,
			pr.images,
			pr.weight,
			pr.price,
			pr.review_count,
			pr.total_rating,
			pr.created_at
		FROM product pr
	`

	params := r.getListFilterParams(filter)

	if len(params) > 0 {
		query += "\n WHERE \n"
		var paramQuery []string
		for key := range params {
			paramQuery = append(paramQuery, fmt.Sprintf("%s = :%s", key, key))
		}

		query += strings.Join(paramQuery, " AND ")
	}

	queryOrderBy := []string{}

	if filter.RatingDescending != nil {
		ratingOrder := "ASC"
		if *filter.RatingDescending {
			ratingOrder = "DESC"
		}

		queryOrderBy = append(queryOrderBy, fmt.Sprintf(`
		CASE
			WHEN review_count = 0 THEN 0
			ELSE total_rating / review_count
		END %s
		`, ratingOrder))
	}

	if filter.CreatedDescending != nil {
		createdOrder := "ASC"
		if *filter.CreatedDescending {
			createdOrder = "DESC"
		}
		queryOrderBy = append(queryOrderBy, fmt.Sprintf(`
			created_at %s
		`, createdOrder))
	}

	if len(queryOrderBy) > 0 {
		joinedOrderBy := strings.Join(queryOrderBy, " , ")
		query = fmt.Sprintf(`
			%s
			ORDER BY %s
		`, query, joinedOrderBy)
	}

	rows, err := tx.NamedQuery(
		query,
		params,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var result entity.ProductSQLX
		if err := rows.StructScan(&result); err != nil {
			return nil, err
		}
		model, err := r.modelToEntity(result)
		if err != nil {
			return nil, err
		}

		models = append(models, *model)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return models, nil
}

func (r *SqlProductRepository) Create(ctx context.Context, product *entity.Product) error {
	tx, err := common.GetTransactionFromContext(ctx, r.db)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO product
			(sku, title, description, category, etalase, images, weight, price, review_count, total_rating, created_at, updated_at)
		VALUES
			(:sku, :title, :description, :category, :etalase, :images, :weight, :price, :review_count, :total_rating, :created_at, :updated_at)
		;
	`

	now := time.Now()

	imagesJSON, err := json.Marshal(product.Images)
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"sku":          product.Sku,
		"title":        product.Title,
		"description":  product.Description,
		"category":     product.Category,
		"etalase":      product.Etalase,
		"images":       imagesJSON,
		"weight":       product.Weight,
		"price":        product.Price,
		"review_count": 0,
		"total_rating": 0,
		"created_at":   now,
		"updated_at":   now,
	}

	_, err = tx.NamedExecContext(
		ctx,
		query,
		params,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *SqlProductRepository) Update(ctx context.Context, product *entity.Product) error {
	tx, err := common.GetTransactionFromContext(ctx, r.db)
	if err != nil {
		return err
	}

	query := `
		UPDATE product
		SET 
			sku = :sku,
			title = :title,
			description = :description,
			category = :category,
			etalase = :etalase,
			images = :images,
			weight = :weight,
			price = :price,
			review_count = :review_count,
			total_rating = :total_rating,
			updated_at = :updated_at
		WHERE
			id = :id
		;
	`
	imagesJSON, err := json.Marshal(product.Images)
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"id":           product.ID,
		"sku":          product.Sku,
		"title":        product.Title,
		"description":  product.Description,
		"category":     product.Category,
		"etalase":      product.Etalase,
		"images":       imagesJSON,
		"weight":       product.Weight,
		"price":        product.Price,
		"review_count": product.ReviewCount,
		"total_rating": product.TotalRating,
		"updated_at":   time.Now(),
	}

	_, err = tx.NamedExecContext(
		ctx,
		query,
		params,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *SqlProductRepository) CreateReview(ctx context.Context, review *entity.Review) error {
	tx, err := common.GetTransactionFromContext(ctx, r.db)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO product_review
			(product_id, rating, review_comment, created_at, updated_at)
		VALUES
			(:product_id, :rating, :review_comment, :created_at, :updated_at)
		;
	`

	now := time.Now()

	params := map[string]interface{}{
		"product_id":     review.ProductID,
		"rating":         review.Rating,
		"review_comment": review.ReviewComment,
		"created_at":     now,
		"updated_at":     now,
	}

	_, err = tx.NamedExecContext(
		ctx,
		query,
		params,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *SqlProductRepository) GetReviewByProductID(ctx context.Context, productID int64) ([]entity.Review, error) {
	tx, err := common.GetTransactionFromContext(ctx, r.db)
	if err != nil {
		return nil, err
	}

	var models = []entity.Review{}

	query := `
		SELECT
			prv.id,
			prv.product_id,
			prv.rating,
			prv.review_comment,
			prv.created_at
		FROM product_review prv
		WHERE prv.product_id = :product_id
		ORDER BY prv.created_at DESC
	`

	params := map[string]interface{}{"product_id": productID}

	rows, err := tx.NamedQuery(
		query,
		params,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var result entity.Review
		if err := rows.StructScan(&result); err != nil {
			return nil, err
		}

		models = append(models, result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return models, nil
}

func (r SqlProductRepository) modelToEntity(model entity.ProductSQLX) (*entity.Product, error) {
	var images []entity.ProductImage
	if err := model.Images.Unmarshal(&images); err != nil {
		// Handle the error
		return nil, err
	}

	var avgRating float64 = 0

	if model.ReviewCount > 0 {
		avgRating = float64(model.TotalRating) / float64(model.ReviewCount)
	}

	return &entity.Product{
		ID:          model.ID,
		Sku:         model.Sku,
		Title:       model.Title,
		Description: model.Description,
		Category:    model.Category,
		Etalase:     model.Etalase,
		Images:      images,
		Weight:      model.Weight,
		Price:       model.Price,
		ReviewCount: model.ReviewCount,
		TotalRating: model.TotalRating,
		AVGRating:   avgRating,
		CreatedAt:   model.CreatedAt,
	}, nil
}

func (r SqlProductRepository) getListFilterParams(filter entity.FilterList) map[string]interface{} {
	var filterAnd = map[string]interface{}{}

	if filter.Sku != "" {
		filterAnd["sku"] = filter.Sku
	}

	if filter.Title != "" {
		filterAnd["title"] = filter.Title
	}

	if filter.Category != "" {
		filterAnd["category"] = filter.Category
	}

	if filter.Etalase != "" {
		filterAnd["etalase"] = filter.Etalase
	}

	return filterAnd
}
