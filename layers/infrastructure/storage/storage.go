package storage

import (
	"fmt"
	"log"
)

type Product struct {
	Id    int
	Name  string
	Price int
}

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

const (
	// StorageServiceName Storage Service Name to access remote procedure
	StorageServiceName = "Storage"
)

var products [10000]Product
var productsIndex = 0

func PrintProductList() {
	for i := 0; i < productsIndex; i++ {
		fmt.Printf("%d %s %d\n", products[i].Id, products[i].Name, products[i].Price)
	}
	fmt.Printf("%d\n", productsIndex)
}

func PrintCurrentItem() {
	//products[productsIndex] = Product{Id: productsIndex, Name: name, Price: price}
	//fmt.Printf("%d %s %d\n", products[i].Id, products[i].Name, products[i].Price)
}

func (s Storage) HelloStorage(args string, args2 string) string {
	log.Println(args, args2)
	return "A Hello message from Storage!\n"
}

func (s Storage) Create(name string, price int) string {
	products[productsIndex] = Product{Id: productsIndex, Name: name, Price: price}
	productsIndex++

	PrintProductList()
	return "Create Function\n"
}

func (s Storage) ReadList() string {
	PrintProductList()

	return "Read Function\n"
}

func (s Storage) ReadItem() string {
	PrintCurrentItem()

	return "Read Function\n"
}

func (s Storage) Update() string {
	//products[productsIndex] = Product{Id: productsIndex, Name: name, Price: price}
	//
	//Product.Id

	return "Update Function\n"
}

func (s Storage) Delete() string {
	//products[productsIndex] = nil
	//productsIndex--
	//fmt.Printf("%d %s %d\n" products[productsIndex])

	return "Delete Function\n"
}

//func (s Storage) AddSlice( []Product) string {
//	products[productsIndex] = Product{Id: productsIndex, Product: product}
//	productsIndex++
//
//	PrintProductList()
//	return "Create Function\n"
//}
