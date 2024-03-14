package cli

import (
	"database/sql"
	"fmt"
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
