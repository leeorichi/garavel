package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var product Product

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=garavel port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.First(&product, 1)
	fmt.Println(product.ID, product.Code, product.Price)
}
