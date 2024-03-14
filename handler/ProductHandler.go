package handler

import (
	"Sport-PairProject/entity"
	"context"
	"database/sql"
	"fmt"
	"time"
)

func InsertNewProduct(db *sql.DB, product entity.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO Products (CategoryID, Name, Description, Price, QuantityInStock) VALUES ($1,$2,$3,$4,$5)"

	_, err := db.ExecContext(ctx, query, product.CategoryID, product.ProductName, product.Description, product.Price, product.QuantityInStock)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func GetCurrentStock(db *sql.DB, product_id int) (current_stock int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "SELECT QuantityInStock FROM Products WHERE ProductID = $1"

	err := db.QueryRowContext(ctx, query, product_id).Scan(&current_stock)
	if err != nil {
		fmt.Println("Failed to get current stock: ", err)
		return
	}

	return current_stock
}

func UpdateProductStock(db *sql.DB, stock, product_id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newStock := GetCurrentStock(db, product_id) + stock

	query := "UPDATE Products SET QuantityInStock = $1 WHERE ProductID = $2"

	_, err := db.ExecContext(ctx, query, newStock, product_id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func DeleteProduct(db *sql.DB, product_id int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "DELETE FROM Products WHERE ProductID = $1"

	res, err := db.ExecContext(ctx, query, product_id)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		return false
	}

	return true
}

func GetAllProductList(db *sql.DB) (products []entity.Product, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var p entity.Product

	query := "SELECT p.ProductID, c.Name, p.Name, p.Description, p.Price, p.QuantityInStock FROM Products p JOIN Categories c ON p.CategoryID = c.CategoryID"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		fmt.Println(err.Error())
		return []entity.Product{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&p.ProductID, &p.CategoryName, &p.ProductName, &p.Description, &p.Price, &p.QuantityInStock)
		if err != nil {
			fmt.Println(err.Error())
			return []entity.Product{}, err
		}
		products = append(products, p)
	}

	return products, nil
}
