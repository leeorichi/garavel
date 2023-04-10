package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"gopkg.in/gookit/color.v1"

	"go.havanlong.com/garavel/routes"
)

func main() {
	godotenv.Load()
	mux := http.NewServeMux()

	color.Green.Println("ahihi đồ chó")
	color.Green.Println("Server was started at http://localhost:3333")

	routers.Init(mux)

	err := http.ListenAndServe(":3333", mux)
	fmt.Println(err)
}
