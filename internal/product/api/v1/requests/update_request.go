package requests

import (
	"net/http"
	"strings"

	"github.com/layardaputra/govtech-catalog-test-project/common"
)

type UpdateRequestParam struct {
	ID          int64
	Sku         string       `json:"sku"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Category    string       `json:"category"`
	Etalase     string       `json:"etalase"`
	Images      []imageParam `json:"images"`
	Weight      float64      `json:"weight"`
	Price       float64      `json:"price"`
}

func (p *UpdateRequestParam) ValidateParam() error {
	p.Sku = strings.TrimSpace(p.Sku)
	p.Title = strings.TrimSpace(p.Title)
	p.Description = strings.TrimSpace(p.Description)
	p.Category = strings.TrimSpace(p.Category)
	p.Etalase = strings.TrimSpace(p.Etalase)
	for i := range p.Images {
		p.Images[i].ImageURL = strings.TrimSpace(p.Images[i].ImageURL)
		p.Images[i].Description = strings.TrimSpace(p.Images[i].Description)
	}

	if p.Weight <= 0 {
		return &common.CustomError{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Invalid Weight",
			Err:        nil,
		}
	}

	if p.Price <= 0 {
		return &common.CustomError{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Invalid Price",
			Err:        nil,
		}
	}

	return nil
}
