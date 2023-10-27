package product_v1

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/layardaputra/govtech-catalog-test-project/common"
	"github.com/layardaputra/govtech-catalog-test-project/internal/product/api/v1/requests"
	"github.com/layardaputra/govtech-catalog-test-project/internal/product/domain/entity"
)

// CreateProduct is a handler that create product.
func (h *HandlerV1) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	productId := chi.URLParam(r, "productID")
	productIdInt, err := strconv.ParseInt(productId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := common.DefaultResponse{
			Message: "Invalid Product ID Given",
		}
		w.Write(resp.ToBytes())
		return
	}

	var data requests.UpdateRequestParam
	// Decode the request body into the struct
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := common.DefaultResponse{
			Message: "Internal Server Error",
		}
		w.Write(resp.ToBytes())
		// Log the error
		log.Printf("Error: %v\nStack Trace:\n%s", r, debug.Stack())
		return
	}
	defer r.Body.Close()

	data.ID = productIdInt

	err = data.ValidateParam()
	if err != nil {
		custErr, ok := err.(*common.CustomError)
		if ok {
			w.WriteHeader(custErr.StatusCode)
			resp := common.DefaultResponse{
				Message: custErr.Message,
			}
			w.Write(resp.ToBytes())
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			resp := common.DefaultResponse{
				Message: "Internal Server Error",
			}
			w.Write(resp.ToBytes())
		}
		// Log the error
		log.Printf("Error: %v\nStack Trace:\n%s", r, debug.Stack())
		return
	}

	var images []entity.ProductImage
	for _, data := range data.Images {
		images = append(images, entity.ProductImage{
			ImageURL:    data.ImageURL,
			Description: data.Description,
		})
	}

	err = common.RunInTrans(ctx, h.DB, func(ctx context.Context) error {
		err := h.ProductService.UpdateProduct(ctx, entity.UpdateProductParams{
			ID:          data.ID,
			Sku:         &data.Sku,
			Title:       &data.Title,
			Description: &data.Description,
			Category:    &data.Category,
			Etalase:     &data.Etalase,
			Images:      images,
			Weight:      &data.Weight,
			Price:       &data.Price,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		custErr, ok := err.(*common.CustomError)
		if ok {
			w.WriteHeader(custErr.StatusCode)
			resp := common.DefaultResponse{
				Message: custErr.Message,
			}
			w.Write(resp.ToBytes())
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			resp := common.DefaultResponse{
				Message: "Internal Server Error",
			}
			w.Write(resp.ToBytes())
		}
		// Log the error
		log.Printf("Error: %v\nStack Trace:\n%s", r, debug.Stack())
		return
	}

	resp := common.DefaultResponse{
		Message: "Success Update Product Data",
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp.ToBytes())
}
