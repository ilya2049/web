package product

func NewService() Service {
	return Service{
		products: make(Products, 0),
	}
}

type Service struct {
	products Products
}

func (s *Service) AddProduct(product Product) {
	s.products = append(s.products, product)
}

func (s Service) GetProducts() Products {
	return s.products
}

func (s Service) GetProductByID(id int) Product {
	for _, product := range s.products {
		if product.ID == id {
			return product
		}
	}

	return Product{}
}

func (s *Service) UpdateProduct(id int, newProduct Product) {
	for i := 0; i < len(s.products); i++ {
		if s.products[i].ID == id {
			s.products[i].Name = newProduct.Name
			s.products[i].Price = newProduct.Price
			break
		}
	}
}
