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

// AddProductReview is a handler that add new product review.
func (h *HandlerV1) AddProductReview(w http.ResponseWriter, r *http.Request) {
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

	var data requests.AddProductReviewRequest
	// Decode the request body into the struct
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := common.DefaultResponse{
			Message: "Internal Server Error",
		}
		w.Write(resp.ToBytes())
		// Log the error
		log.Printf("Error: %v\nStack Trace:\n%s", err, debug.Stack())
		return
	}
	defer r.Body.Close()

	data.ProductID = productIdInt

	err = data.ValidateRequest()
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
		log.Printf("Error: %v\nStack Trace:\n%s", err, debug.Stack())
		return
	}

	err = common.RunInTrans(ctx, h.DB, func(ctx context.Context) error {
		err := h.ProductService.AddReviewProduct(ctx, entity.CreateReviewParam{
			ProductID:     data.ProductID,
			Rating:        data.Rating,
			ReviewComment: data.ReviewComment,
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
		log.Printf("Error: %v\nStack Trace:\n%s", err, debug.Stack())
		return
	}

	resp := common.DefaultResponse{
		Message: "Success Add New Product Review",
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp.ToBytes())
}
