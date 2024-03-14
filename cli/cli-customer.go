package cli

import (
	"Sport-PairProject/handler"
	"bufio"
	"database/sql"
	"fmt"
	"os"
)

func CustomerCLI(db *sql.DB, user_id int) {
	var opt int

	for {

		ShowMenuCustomer()

		fmt.Print("Input > ")
		_, err := fmt.Scanln(&opt)
		if err != nil {
			fmt.Println("Input proper value")
			continue
		}

		switch opt {
		case 1:

		case 2:
    for {
			var userID int
			var productID int
			var quantity int
			var selectedMenu string
			fmt.Print("Enter user id :")
			fmt.Scanln(&userID)
			fmt.Print("Enter product id:")
			fmt.Scanln(&productID)
			fmt.Print("Enter quantity :")
			fmt.Scanln(&quantity)
			err := handler.AddItemToCart(db, userID, productID, quantity)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print("Do you want to add another item ? (Y/N):")
			fmt.Scanln(&selectedMenu)
			if selectedMenu == "Y" {
				continue
			} else if selectedMenu == "N" {
				fmt.Println("Check Out Process...")
				handler.CheckoutCartItems(db, userID)
			} else {
				fmt.Println("Enter Y/N only")
				continue
			}
			fmt.Println("Enter your payment method (Cash, Member Balance, Debit card, Credit Card) :")
			paymentMethod := bufio.NewScanner(os.Stdin)
			paymentMethod.Scan()
			err = paymentMethod.Err()
			if err != nil {
				fmt.Println(err)
			}
			getPaymentMethod := paymentMethod.Text()
			err = handler.ProcessPayment(db, userID, getPaymentMethod)
			if err != nil {
				fmt.Println(err)
			}
			return

			}

		case 3:
			return
		}
	}
}

func ShowMenuCustomer() {
	fmt.Println("=== CUSTOMER MENU ====")
	fmt.Println("1. Show all list of products")
	fmt.Println("2. Add product to cart")
	fmt.Println("3. Logout")
}
