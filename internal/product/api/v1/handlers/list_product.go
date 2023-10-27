package product_v1

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/layardaputra/govtech-catalog-test-project/common"
	"github.com/layardaputra/govtech-catalog-test-project/internal/product/api/v1/requests"
	"github.com/layardaputra/govtech-catalog-test-project/internal/product/domain/entity"
)

// ListProduct is a handler that list products.
func (h *HandlerV1) ListProduct(w http.ResponseWriter, r *http.Request) {
	queryRequest := requests.TransformQueryParam(r.URL.Query())

	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	res, err := h.ProductService.ListProducts(ctx, entity.ListProductParams{
		Sku:         queryRequest.Sku,
		Title:       queryRequest.Title,
		Category:    queryRequest.Category,
		Etalase:     queryRequest.Etalase,
		SortCreated: queryRequest.SortCreated,
		SortRating:  queryRequest.SortRating,
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

	resp := common.ResponseSuccessWithData{
		Message: "Success Get List Product Datas",
		Data:    res,
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp.ToBytes())

}
