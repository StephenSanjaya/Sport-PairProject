package cli

import (
	"Sport-PairProject/entity"
	"Sport-PairProject/handler"
	"database/sql"
	"fmt"
)

func AdminCLI(db *sql.DB, user_id int) {

	var opt int

	for {

		ShowMenuAdmin()

		fmt.Print("Input > ")
		_, err := fmt.Scanln(&opt)
		if err != nil {
			fmt.Println("Input proper value")
			continue
		}

		switch opt {
		case 1:
			ShowAllProductList(db)
		case 2:
			AddNewProduct(db)
		case 3:
			AddMoreStockToProduct(db)
		case 4:
			RemoveProduct(db)
		case 5:
			GenerateUserReport(db)
		case 6:
			return
		}
	}
}

func ShowMenuAdmin() {
	fmt.Println("=== ADMIN MENU ====")
	fmt.Println("1. Show all list of products")
	fmt.Println("2. Add new product")
	fmt.Println("3. Increase quantity product")
	fmt.Println("4. Remove product from menu")
	fmt.Println("================================")
	fmt.Println("5. Generate user report")
	fmt.Println("6. Logout")
}

func GenerateUserReport(db *sql.DB) {
	user_report, err := handler.GetUserReport(db)
	if err != nil {
		fmt.Println("Failed to get user report", err)
		return
	}
	for _, v := range user_report {
		fmt.Println(v)
	}
}

func ShowAllProductList(db *sql.DB) {
	products, err := handler.GetAllProductList(db)
	if err != nil {
		fmt.Println("Failed to show all list of product: ", err)
		return
	}

	fmt.Println("PRODUCT LIST:")
	for _, v := range products {
		fmt.Println()
		fmt.Println("ID: ", v.ProductID)
		fmt.Println("Category: ", v.CategoryName)
		fmt.Println("Product: ", v.ProductName)
		fmt.Println("Description: ", v.Description)
		fmt.Println("Price: ", v.Price)
		fmt.Println("Quantity: ", v.QuantityInStock)
	}
}

func AddNewProduct(db *sql.DB) {
	var category_id, qty int
	var product_name, description string
	var price float64

	fmt.Print("Insert Category ID: ")
	_, err := fmt.Scanln(&category_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print("Insert product name: ")
	_, err = fmt.Scanln(&product_name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print("Insert product description: ")
	_, err = fmt.Scanln(&description)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print("Insert product price: ")
	_, err = fmt.Scanln(&price)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print("Insert product quantity: ")
	_, err = fmt.Scanln(&qty)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var new_product = entity.Product{
		CategoryID:      category_id,
		ProductName:     product_name,
		Description:     description,
		Price:           price,
		QuantityInStock: qty,
	}

	err = handler.InsertNewProduct(db, new_product)
	if err != nil {
		fmt.Println("Failed to insert new product: ", err)
		return
	}

	fmt.Println("Successfully add new product")
}

func AddMoreStockToProduct(db *sql.DB) {
	var product_id, qty int
	fmt.Println("Input Product ID: ")
	_, err := fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Input quantity: ")
	_, err = fmt.Scanln(&qty)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = handler.UpdateProductStock(db, qty, product_id)
	if err != nil {
		fmt.Println("Failed to update product quantity: ", err)
		return
	}

	fmt.Println("Success to increase the product quantity")
}

func RemoveProduct(db *sql.DB) {
	var product_id int

	fmt.Print("Input Product ID you want to delete: ")
	_, err := fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = handler.DeleteProduct(db, product_id)
	if err != nil {
		fmt.Println("Failed to delete product: ", err)
		return
	}
}
