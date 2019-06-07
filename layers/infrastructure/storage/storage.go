package storage

import "fmt"

type Product struct {
	Id    int
	Name  string
	Price float64
}

type Storage struct {
	products      []Product
	productsIndex int
}

func NewStorage() *Storage {
	return &Storage{
		products:      make([]Product, 10000),
		productsIndex: 0,
	}
}

const (
	// StorageServiceName Storage Service Name to access remote procedure
	StorageServiceName = "Storage"
)

func (s *Storage) PrintProductList() string {

	response := "\n\nProduct List: \n"
	for i := 0; i <= s.productsIndex-1; i++ {
		response += "Product " + fmt.Sprintf("%d", s.products[i].Id) + ": " + s.products[i].Name + " R$" + fmt.Sprintf("%.2f", s.products[i].Price) + "\n"
	}
	return response
}

func (s *Storage) Create(name string, price float64) string {

	s.products[s.productsIndex] = Product{Id: s.productsIndex, Name: name, Price: price}
	s.productsIndex++

	// return s.PrintProductList() + "Create Function Complete\n\n"
	return "Create Function Complete\n"

}

func (s *Storage) ReadList(name string, price float64) string {
	return s.PrintProductList() + "Read List Function Complete\n"
}

func (s *Storage) ReadItem(name string) string {

	response := ""
	for _, p := range s.products {
		if p.Name == name {
			response += "Product " + fmt.Sprintf("%d", p.Id) + ": " + p.Name + "Price: " + fmt.Sprintf("%.2f", p.Price) + "\n"
		}
	}
	if response == "" {
		return "Read Item Function Failed To Find Item\n"
	} else {
		response += "Read Item Function Complete\n"
		return response
	}
}

func (s *Storage) Update(name string, price float64) string {

	for i, p := range s.products {
		if p.Name == name {
			s.products[i].Price = price
		}
	}
	return "Update Function Complete\n"
}

func (s *Storage) Delete() string {

	s.products[s.productsIndex] = Product{0, "", 0}
	s.productsIndex--
	return "Delete Function Complete\n"
}
