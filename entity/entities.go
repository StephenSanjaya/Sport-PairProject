package entity

type User struct {
	UserID                                   int
	Balance                                  float64
	Username, Password, Email, Address, Role string
}

type Category struct {
	CategoryID int
	Name       string
}

type Product struct {
	ProductID, CategoryID, QuantityInStock int
	CategoryName, ProductName, Description string
	Price                                  float64
}

type Cart struct {
	CartID, UserID, ProductID, Quantity int
	SubTotal                            float64
}

type Order struct {
	OrderID, UserID, ProductID, Quantity int
	SubTotal                             float64
	OrderDate, Status, PaymentMethod     string
}

type UserReport struct {
	UserID                   int
	Username, Email, Address string
	Balance                  float64
}
