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

// For Report
type UserTransaction struct {
	Username                 string
	TransactionCount         int
	TotalQuantity            int
	TotalAmount              float64
	AverageAmountTransaction float64
	LastTransactionDate      string
}

type CategorySales struct {
	CategoryName   string
	TotalQtySold   int
	TotalAmount    float64
}

type ProductSales struct {
	ProductName           string
	ProductCategory       string
	TransactionCount      int
	TotalQtySold          int
	TotalAmount           float64
	AvgAmountTransaction  float64
}

type OrderStatus struct {
	Status            string
	TransactionCount  int
	TotalAmount       float64
}