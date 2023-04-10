package routers

import (
	"net/http"

	controllers "go.havanlong.com/garavel/app/Http/Controllers"
)

// m.HandleFunc("<patten>", <controller.method()>)

func Init(m *http.ServeMux) *http.ServeMux {
	m.HandleFunc("/", controllers.Index)
	return m
}
