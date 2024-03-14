package entity

type Products struct {
	ProductID, CategoryID, QuantityInStock int
	CategoryName, ProductName, Description string
	Price                                  float64
}
