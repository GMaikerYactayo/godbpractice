package main

import (
	"fmt"
	"github.com/GMaikerYactayo/godbpractice/pkg/product"
	"github.com/GMaikerYactayo/godbpractice/storage"
	"log"
)

func main() {
	storage.NewMySqlDB()

	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		Name:         "Umbrella",
		Observations: "Plastic",
		Price:        20,
	}
	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.Create: %v", err)
	}

	fmt.Printf("%+v\n", m)

}
