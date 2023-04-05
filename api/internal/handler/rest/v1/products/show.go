package products

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/leeorichi/getgo/api/internal/httpserver"
)

// customRespond

// type CustomRespond struct {
// 	Status  bool   `json:"status"`
// 	Message string `json:"message"`
// }

// Create creates new product
func (h Handler) Show() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		eid := chi.URLParam(r, "eid")
		if p, err := h.productCtrl.Show(ctx, eid); err != nil {
			httpserver.RespondJSON(w, CustomRespond{
				Status:  false,
				Message: "Product was not exist",
			})
		} else {
			httpserver.RespondJSON(w, p)
		}
		return nil
	})
}
