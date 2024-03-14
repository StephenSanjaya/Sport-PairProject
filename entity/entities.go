package entity

type User struct {
	UserID   int
	Balance  float64
	Username, Password, Email, Address, Role string
}

type Category struct {
	CategoryID int
	Name       string
}

type Product struct {
	ProductID, CategoryID, QuantityInStock int
	Price                                 float64
	Name, Description                     string
}

type Cart struct {
	CartID, UserID, ProductID, Quantity int
	SubTotal                           float64
}

type Order struct {
	OrderID, UserID, ProductID, Quantity int
	SubTotal                            float64
	OrderDate, Status, PaymentMethod    string
}

