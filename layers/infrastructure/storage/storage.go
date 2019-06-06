package storage

import "strconv"

type Product struct {
	Id    int
	Name  string
	Price int
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

	response := "\n"
	for i := 0; i <= s.productsIndex; i++ {
		response += "Product: " + s.products[i].Name + "Price: " + strconv.Itoa(s.products[i].Price) + "\n"
	}
	return response
}

// func (s *Storage) PrintCurrentItem() {
// 	//products[productsIndex] = Product{Id: productsIndex, Name: name, Price: price}
// 	//fmt.Printf("%d %s %d\n", products[i].Id, products[i].Name, products[i].Price)
// }

// func (s *Storage) HelloStorage(args string, args2 string) string {
// 	log.Println(args, args2)
// 	return "A Hello message from Storage!\n"
// }

func (s *Storage) Create(name string, price int) string {

	s.products[s.productsIndex] = Product{Id: s.productsIndex, Name: name, Price: price}
	s.productsIndex++

	s.PrintProductList()
	return "Create Function Complete\n"
}

func (s *Storage) ReadList() string {
	s.PrintProductList()
	return "Read List Function Complete\n"
}

func (s *Storage) ReadItem(name string) string {

	response := ""
	for _, p := range s.products {
		if p.Name == name {
			response += "Product: " + p.Name + "Price: " + strconv.Itoa(p.Price) + "\n"
		}
	}
	if response == "" {
		return "Read Item Function Failed To Find Item\n"
	} else {
		response += "Read Item Function Complete\n"
		return response
	}
}

func (s *Storage) Update(name string, price int) string {

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

//func (s Storage) AddSlice( []Product) string {
//	products[productsIndex] = Product{Id: productsIndex, Product: product}
//	productsIndex++
//
//	PrintProductList()
//	return "Create Function\n"
//}
