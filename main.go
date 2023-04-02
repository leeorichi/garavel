// main.go
package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"gopkg.in/gookit/color.v1"
	"go-tidypj/src/httpe"
)

func main() {
	godotenv.Load()
	color.Red.Println("Red color")
	m := httpe.Getgo()
	fmt.Println(m)
}
