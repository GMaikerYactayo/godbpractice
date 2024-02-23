package main

import (
	"fmt"
	"github.com/GMaikerYactayo/godbpractice/pkg/product"
	"github.com/GMaikerYactayo/godbpractice/storage"
	"log"
	"time"
)

func main() {
	driver := storage.MySQL

	storage.New(driver)

	myStorage, err := storage.DAOProduct(driver)
	if err != nil {
		log.Fatalf("DAOProduct: %v", err)
	}
	serviceProduct := product.NewService(myStorage)

	ms, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}
	for _, m := range ms {
		fmt.Printf("ID: %d, Name: %s, Observations: %s, Price: %d, CreateAt: %s, UpdatedAt: %s\n",
			m.ID, m.Name, m.Observations, m.Price, m.CreateAt.Format(time.RFC822), m.UpdateAt.Format(time.RFC822))
	}
}
