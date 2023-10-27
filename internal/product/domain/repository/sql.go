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
			pr.price
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

	var models []entity.Product

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
			pr.price
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

	orderBy := "ASC"
	if filter.IsDescending {
		orderBy = "DESC"
	}

	query = fmt.Sprintf("%s \nORDER BY created_at %s", query, orderBy)

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
			(sku, title, description, category, etalase, images, weight, price, created_at, updated_at)
		VALUES
			(:sku, :title, :description, :category, :etalase, :images, :weight, :price, :created_at, :updated_at)
		;
	`

	now := time.Now()

	imagesJSON, err := json.Marshal(product.Images)
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"sku":         product.Sku,
		"title":       product.Title,
		"description": product.Description,
		"category":    product.Category,
		"etalase":     product.Etalase,
		"images":      imagesJSON,
		"weight":      product.Weight,
		"price":       product.Price,
		"created_at":  now,
		"updated_at":  now,
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
		"id":          product.ID,
		"sku":         product.Sku,
		"title":       product.Title,
		"description": product.Description,
		"category":    product.Category,
		"etalase":     product.Etalase,
		"images":      imagesJSON,
		"weight":      product.Weight,
		"price":       product.Price,
		"updated_at":  time.Now(),
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

func (r SqlProductRepository) modelToEntity(model entity.ProductSQLX) (*entity.Product, error) {
	var images []entity.ProductImage
	if err := model.Images.Unmarshal(&images); err != nil {
		// Handle the error
		return nil, err
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
