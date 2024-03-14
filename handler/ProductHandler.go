package handler

import (
	"Sport-PairProject/entity"
	"context"
	"database/sql"
	"fmt"
	"time"
)

func InsertNewProduct(db *sql.DB, product entity.Products) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO Products (CategoryID, Name, Description, Price, QuantityInStock) VALUES (?,?,?,?)"

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

	query := "SELECT QuantityInStock FROM Products WHERE ProductID = ?"

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

	query := "UPDATE Products WHERE QuantityInStock = ?"

	_, err := db.ExecContext(ctx, query, newStock)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func DeleteProduct(db *sql.DB, product_id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "DELETE Products WHERE ProductID = ?"

	_, err := db.ExecContext(ctx, query, product_id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func GetAllProductList(db *sql.DB) (products []entity.Products, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var p entity.Products

	query := "SELECT p.ProductID, c.Name, p.Name, p.Description, p.Price, p.QuantityInStock FROM Products p JOIN Categoryies c ON p.CategoryID = c.CategoryID"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		fmt.Println(err.Error())
		return []entity.Products{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&p.ProductID, &p.CategoryName, &p.ProductName, &p.Description, &p.Price, &p.QuantityInStock)
		if err != nil {
			fmt.Println(err.Error())
			return []entity.Products{}, err
		}
		products = append(products, p)
	}

	return products, nil
}
