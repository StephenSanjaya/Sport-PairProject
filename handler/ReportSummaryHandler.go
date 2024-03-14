package handler

import (
	"Sport-PairProject/entity"
	"database/sql"
	"fmt"
)

func GetUserTransactions(db *sql.DB) ([]entity.UserTransaction, error) {
	rows, err := db.Query(`
		SELECT u.Username AS user, COUNT(o.OrderID) AS TRANSACTION_COUNT, SUM(o.Quantity) AS TOTAL_QUANTITY, 
		SUM(o.Subtotal) AS TOTAL_AMOUNT, AVG(o.Subtotal) AS AVERAGE_AMOUNT_TRANSACTION, MAX(o.OrderDate) AS LAST_TRANSACTION_DATE
		FROM Users u
		JOIN Orders o
		ON u.userID = o.UserID
		GROUP BY u.Username;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []entity.UserTransaction
	for rows.Next() {
		var t entity.UserTransaction
		err := rows.Scan(&t.Username, &t.TransactionCount, &t.TotalQuantity, &t.TotalAmount, &t.AverageAmountTransaction, &t.LastTransactionDate)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func GetCategorySales(db *sql.DB) ([]entity.CategorySales, error) {
	rows, err := db.Query(`
		SELECT c.Name AS CATEGORY_NAME, SUM(o.Quantity) AS TOTAL_QTY_SOLD, SUM(o.SubTotal) AS TOTAL_AMOUNT
		FROM Categories c
		JOIN Products p
		ON c.CategoryID = p.CategoryID
		JOIN Orders o
		ON p.ProductID = o.ProductID
		GROUP BY CATEGORY_NAME
		ORDER BY TOTAL_QTY_SOLD DESC;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []entity.CategorySales
	for rows.Next() {
		var s entity.CategorySales
		err := rows.Scan(&s.CategoryName, &s.TotalQtySold, &s.TotalAmount)
		if err != nil {
			return nil, err
		}
		sales = append(sales, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sales, nil
}

func GetProductSales(db *sql.DB) ([]entity.ProductSales, error) {
	rows, err := db.Query(`
		SELECT p.Name AS PRODUCT_NAME, c.Name AS PRODUCT_CATEGORY, COUNT(o.OrderID) AS TRANSACTION_COUNT, 
		SUM(o.Quantity) AS TOTAL_QTY_SOLD, SUM(o.Subtotal) AS TOTAL_AMOUNT, AVG(o.Subtotal) AS AVG_AMOUNT_TRANSACTION
		FROM Products p
		JOIN Categories c
		ON p.CategoryID = c.CategoryID
		JOIN Orders o
		ON p.ProductID = o.ProductID
		GROUP BY p.Name, c.Name
		ORDER BY TOTAL_QTY_SOLD DESC;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []entity.ProductSales
	for rows.Next() {
		var s entity.ProductSales
		err := rows.Scan(&s.ProductName, &s.ProductCategory, &s.TransactionCount, &s.TotalQtySold, &s.TotalAmount, &s.AvgAmountTransaction)
		if err != nil {
			return nil, err
		}
		sales = append(sales, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sales, nil
}

func GetOrderStatus(db *sql.DB) ([]entity.OrderStatus, error) {
	rows, err := db.Query(`
		SELECT status, COUNT(*) AS TRANSACTION_COUNT, SUM(Subtotal) AS TOTAL_AMOUNT
		FROM Orders
		WHERE Status IN ('Completed', 'Pending', 'Cancelled')
		GROUP BY status
		ORDER BY TRANSACTION_COUNT DESC;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []entity.OrderStatus
	for rows.Next() {
		var s entity.OrderStatus
		err := rows.Scan(&s.Status, &s.TransactionCount, &s.TotalAmount)
		if err != nil {
			return nil, err
		}
		statuses = append(statuses, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return statuses, nil
}

func PrintTransactions(transactions []entity.UserTransaction) {
	fmt.Println("User Transactions Report:")
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("|%-10s | %-20s | %-20s | %-20s | %-30s | %-25s |\n", "User Name", "Transaction Count", "Total Quantity", "Total Amount", "Average Amount Transaction", "Last Transaction Date")
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------")
	for _, t := range transactions {
		fmt.Printf("|%-10s | %-20d | %-20d | %-20.2f | %-30.2f | %-25s |\n", t.Username, t.TransactionCount, t.TotalQuantity, t.TotalAmount, t.AverageAmountTransaction, t.LastTransactionDate)
	}
}

func PrintCategorySales(sales []entity.CategorySales) {
	fmt.Println("Category Sales Report:")
	fmt.Println("---------------------------------------------------------------------")
	fmt.Printf("|%-20s | %-20s | %-20s |\n", "Category Name", "Total Qty Sold", "Total Amount")
	fmt.Println("---------------------------------------------------------------------")
	for _, s := range sales {
		fmt.Printf("|%-20s | %-20d | %-20.2f |\n", s.CategoryName, s.TotalQtySold, s.TotalAmount)
	}
}

func PrintProductSales(sales []entity.ProductSales) {
	fmt.Println("Product Sales Report:")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("|%-30s | %-20s | %-20s | %-20s | %-20s | %-30s |\n", "Product Name", "Product Category", "Transaction Count", "Total Qty Sold", "Total Amount", "Avg Amount Transaction")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------------------")
	for _, s := range sales {
		fmt.Printf("|%-30s | %-20s | %-20d | %-20d | %-20.2f | %-30.2f |\n", s.ProductName, s.ProductCategory, s.TransactionCount, s.TotalQtySold, s.TotalAmount, s.AvgAmountTransaction)
	}
}

func PrintOrderStatus(statuses []entity.OrderStatus) {
	fmt.Println("Order Status Report:")
	fmt.Println("---------------------------------------------------------------------")
	fmt.Printf("|%-20s | %-20s | %-20s |\n", "Status", "Transaction Count", "Total Amount")
	fmt.Println("---------------------------------------------------------------------")
	for _, s := range statuses {
		fmt.Printf("|%-20s | %-20d | %-20.2f |\n", s.Status, s.TransactionCount, s.TotalAmount)
	}
}