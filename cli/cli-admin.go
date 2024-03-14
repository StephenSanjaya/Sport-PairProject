package cli

import (
	"Sport-PairProject/entity"
	"Sport-PairProject/handler"
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func AdminCLI(db *sql.DB, user_id int) {

	var opt int
	stdin := bufio.NewReader(os.Stdin)

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
			AddNewProduct(db, stdin)
		case 3:
			AddMoreStockToProduct(db)
		case 4:
			RemoveProduct(db)
		case 5:
			GenerateUserReport(db)
		case 6:
			GenerateStockReport(db)
		case 7:
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
	fmt.Println("=========== REPORTING ============")
	fmt.Println("5. Generate user report")
	fmt.Println("6. Generate stock report")
	fmt.Println("7. Logout")
}

func GenerateStockReport(db *sql.DB) {
	stock_report, err := handler.GetStockReport(db)
	if err != nil {
		fmt.Println("Failed to get stock report", err)
		return
	}

	fmt.Println("==== STOCK REPORT ====")
	for _, v := range stock_report {
		fmt.Println("Product ID: ", v.ProductID)
		fmt.Println("Product Name: ", v.ProductName)
		fmt.Println("Product Price: ", v.Price)
		fmt.Println("Product Quantity: ", v.QuantityInStock)
		fmt.Println()
	}
}

func GenerateUserReport(db *sql.DB) {
	user_report, err := handler.GetUserReport(db)
	if err != nil {
		fmt.Println("Failed to get user report", err)
		return
	}

	fmt.Println("==== USER REPORT ====")
	for _, v := range user_report {
		fmt.Println("User ID: ", v.UserID)
		fmt.Println("User Name: ", v.Username)
		fmt.Printf("User Balance: %.2f\n", v.Balance)
		fmt.Println("User Address: ", v.Address)
		fmt.Println("User Email: ", v.Email)
		fmt.Println()
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
		fmt.Println("ID: ", v.ProductID)
		fmt.Println("Category: ", v.CategoryName)
		fmt.Println("Product: ", v.ProductName)
		fmt.Println("Description: ", v.Description)
		fmt.Println("Price: ", v.Price)
		fmt.Println("Quantity: ", v.QuantityInStock)
		fmt.Println()
	}
}

func AddNewProduct(db *sql.DB, stdin *bufio.Reader) {
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
	product_name, err = stdin.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		stdin.ReadString('\n')
		return
	}
	product_name = strings.TrimSpace(product_name)

	fmt.Print("Insert product description: ")
	description, err = stdin.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	description = strings.TrimSpace(description)

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
	fmt.Println()
}

func AddMoreStockToProduct(db *sql.DB) {
	var product_id, qty int
	fmt.Print("Input Product ID: ")
	_, err := fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print("Input quantity: ")
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
	fmt.Println()
}

func RemoveProduct(db *sql.DB) {
	var product_id int

	fmt.Print("Input Product ID you want to delete: ")
	_, err := fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	boolean := handler.DeleteProduct(db, product_id)
	if !boolean {
		fmt.Println("Failed to delete product: product id not found")
		return
	}

	fmt.Println("Successfully remove the product")
	fmt.Println()
}
