package storage

import(
	"log"
	"fmt"
)

type Product struct {
	Id 		int
	Name    string
	Price   int 
}

type Storage struct{}

const (
	// StorageServiceName Storage Service Name to access remote procedure
	StorageServiceName = "Storage"
)

func NewStorage() *Storage {
	return &Storage{}
}

func PrintProductList() {
	for i := 0; i < productsIndex; i++ {
		fmt.Printf("%d %s %d\n", products[i].Id, products[i].Name, products[i].Price)
	}
	fmt.Printf("%d\n", productsIndex)
}

var products [4]Product
var productsIndex = 0

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

func (s Storage) Read() string {
	return "Read Function\n"
}

func (s Storage) Update() string {
	return "Update Function\n"
}

func (s Storage) Delete() string {
	return "Delete Function\n"
}
