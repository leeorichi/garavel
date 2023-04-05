package products

import (
	"net/http"

	"github.com/leeorichi/getgo/api/internal/httpserver"
)

// customRespond

// type CustomRespond struct {
// 	Status  bool   `json:"status"`
// 	Message string `json:"message"`
// }

// Create creates new product
func (h Handler) All() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		if p, err := h.productCtrl.All(ctx); err != nil {
			httpserver.RespondJSON(w, CustomRespond{
				Status:  false,
				Message: "Product is empty",
			})
		} else {
			httpserver.RespondJSON(w, p)
		}
		return nil
	})
}
