package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gookit/color.v1"

	routers "go.havanlong.com/garavel/routes"
)

var env = "dev"

func main() {

	switch env {
	case "qc":
		godotenv.Load(".env.qc")

	case "dev":
		godotenv.Load(".env.local")

	}

	dd := os.Getenv("APP_NAME")

	fmt.Println(dd)

	mux := http.NewServeMux()

	color.Green.Println("ahihi đồ chó")
	color.Green.Println("Server was started at http://localhost:3333")

	routers.Init(mux)

	err := http.ListenAndServe(":3333", mux)
	fmt.Println(err)
}
