package products

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/leeorichi/getgo/api/internal/httpserver"
)

// customRespond
type CustomRespond struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// Create creates new product
func (h Handler) Delete() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		eid := chi.URLParam(r, "eid")
		m := ""

		fmt.Println("==>: ")
		fmt.Println(eid)
		fmt.Println("<==")

		stt, err := h.productCtrl.Delete(ctx, eid)
		if err != nil {
			m = err.Error()
		} else {
			m = "Product was deleted"
		}
		// m := map[bool]string{err != nil: err.Error(), err == nil: "Product was deleted."}[err == nil]
		rp := CustomRespond{
			Status:  stt,
			Message: m,
		}
		httpserver.RespondJSON(w, rp)
		return nil
	})
}
