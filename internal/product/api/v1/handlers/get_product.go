package product_v1

import (
	"log"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/layardaputra/govtech-catalog-test-project/common"
)

// GetProduct is a handler that get product by id.
func (h *HandlerV1) GetProduct(w http.ResponseWriter, r *http.Request) {
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

	res, err := h.ProductService.GetProductByID(ctx, productIdInt)
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

	resp := common.ResponseSuccessWithData{
		Message: "Success Get Product Data",
		Data:    res,
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp.ToBytes())
}
