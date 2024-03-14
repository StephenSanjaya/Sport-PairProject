package handler

import (
	"database/sql"
	"fmt"
	"time"
)

func AddItemToCart(db *sql.DB, userID, productID, quantity int) error {
	// Retrieve the price for the product
	var price float64
	err := db.QueryRow("SELECT Price FROM Products WHERE ProductID = $1", productID).Scan(&price)
	if err != nil {
		return err
	}

	// Calculate SubTotal in Go
	subTotal := price * float64(quantity)

	// Add item to Cart with the calculated SubTotal
	_, err = db.Exec("INSERT INTO Carts (UserID, ProductID, Quantity, SubTotal) VALUES ($1, $2, $3, $4)",
		userID, productID, quantity, subTotal)
	if err != nil {
		return err
	}
	return nil
}

func CheckoutCartItems(db *sql.DB, userID int) error {
	// Get current date
	currentDate := time.Now()

	// Retrieve all cart items for the user and insert them into Orders directly
	rows, err := db.Query(`
			SELECT c.ProductID, c.Quantity, p.Price, p.QuantityInStock
			FROM Carts c
			JOIN Products p ON c.ProductID = p.ProductID
			WHERE c.UserID = $1
	`, userID)
	if err != nil {
			return fmt.Errorf("error retrieving cart items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
			var productID, quantity, quantityInStock int
			var price float64
			err = rows.Scan(&productID, &quantity, &price, &quantityInStock)
			if err != nil {
					return fmt.Errorf("error scanning row: %w", err)
			}

			// Adjust quantity based on stock availability
			if quantity > quantityInStock {
					fmt.Printf("The quantity for ProductID %d has been adjusted to %d due to stock limitations.\n", productID, quantityInStock)
					quantity = quantityInStock
					// add return adjust by customer
			}

			subTotal := price * float64(quantity)

			// Insert directly into Orders without a prepared statement
			_, err = db.Exec(`
					INSERT INTO Orders (UserID, OrderDate, ProductID, Quantity, SubTotal, Status, PaymentMethod) 
					VALUES ($1, $2, $3, $4, $5, 'Pending', 'Cash')
			`, userID, currentDate, productID, quantity, subTotal)
			if err != nil {
					return fmt.Errorf("error inserting order: %w", err)
			}
	}

	// Delete cart items after moving to Orders
	_, err = db.Exec("DELETE FROM Carts WHERE UserID = $1", userID)
	if err != nil {
			return fmt.Errorf("error deleting cart items: %w", err)
	}

	return nil
}

func ProcessPayment(db *sql.DB, userID int, paymentMethod string) error {
	var totalAmount float64
	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Retrieve each pending order for the user
	rows, err := tx.Query(`
		SELECT o.OrderID, o.ProductID, o.Quantity
		FROM Orders o
		WHERE o.UserID = $1 AND o.Status = 'Pending'
	`, userID)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var orderID, productID, quantity int
		err = rows.Scan(&orderID, &productID, &quantity)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Retrieve the total amount of pending orders for the user
	err = tx.QueryRow(`
		SELECT SUM(SubTotal) FROM Orders 
		WHERE UserID = $1 AND Status = 'Pending'
	`, userID).Scan(&totalAmount)
	if err != nil {
		tx.Rollback()
		return err
	}

	// If Member Balance is chosen, check if the user's balance is sufficient
	if paymentMethod == "Member Balance" {
		var balance float64
		err = tx.QueryRow("SELECT Balance FROM Users WHERE UserID = $1", userID).Scan(&balance)
		if err != nil {
			tx.Rollback()
			return err
		}

		// If balance is insufficient, ask the customer to choose another payment method
		if totalAmount > balance {
			tx.Rollback()
			return fmt.Errorf("insufficient balance, please choose another payment method")
		}

		// Deduct the total amount from the user's balance
		_, err = tx.Exec("UPDATE Users SET Balance = Balance - $2 WHERE UserID = $1", userID, totalAmount)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Update the order status to 'Completed' for all valid orders
	_, err = tx.Exec(`
		UPDATE Orders SET Status = 'Completed', PaymentMethod = $2 
		WHERE UserID = $1 AND Status = 'Pending'
	`, userID, paymentMethod)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Decrease the product quantity in the Products table for completed orders
	_, err = tx.Exec(`
		UPDATE Products SET QuantityInStock = QuantityInStock - o.Quantity
		FROM Orders o
		WHERE o.ProductID = Products.ProductID AND o.UserID = $1 AND o.Status = 'Completed'
	`, userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}