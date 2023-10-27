package requests

import (
	"net/url"
	"strings"
)

type ListQueryParams struct {
	Sku      string `json:"sku"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Etalase  string `json:"etalase"`
	Sort     string `json:"sort"`
}

func TransformQueryParam(queryParams url.Values) ListQueryParams {
	var query ListQueryParams

	for key, values := range queryParams {
		switch key {
		case "sku":
			query.Sku = strings.TrimSpace(values[0])
		case "category":
			query.Category = strings.TrimSpace(values[0])
		case "title":
			query.Title = strings.TrimSpace(values[0])
		case "etalase":
			query.Etalase = strings.TrimSpace(values[0])
		case "sort":
			query.Sort = strings.TrimSpace(values[0])
		}
	}

	return query
}
