package cli

import (
	"Sport-PairProject/config"
	"Sport-PairProject/entity"
	"Sport-PairProject/handler"
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func RunApplication() {

	db, err := config.GetConnection()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Menu:")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")

		fmt.Print("Choose an option: ")
		optionStr, _ := reader.ReadString('\n')
		option := strings.TrimSpace(optionStr)

		switch option {
		case "1":
			Register(db, reader)
		case "2":
			role, user_id := Login(db, reader)
			if role == "Customer" {
				CustomerCLI(db, user_id)
			} else if role == "Admin" {
				AdminCLI(db, user_id)
			}
		case "3":
			fmt.Println("Exiting program...")
			os.Exit(1)
		default:
			fmt.Println("Invalid option. Please choose again.")
		}
	}

}

func Register(db *sql.DB, reader *bufio.Reader) {
	fmt.Print("\nEnter username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input proper value", err.Error())
		return
	}
	username = strings.TrimSpace(username)

	fmt.Print("Enter email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input proper value", err.Error())
		return
	}
	email = strings.TrimSpace(email)

	fmt.Print("Enter password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input proper value", err.Error())
		return
	}
	password = strings.TrimSpace(password)

	fmt.Print("Enter address: ")
	address, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input proper value", err.Error())
		return
	}
	address = strings.TrimSpace(address)

	newUser := entity.User{
		Username: username,
		Email:    email,
		Password: password,
		Address:  address,
		Balance:  200000,
	}

	err = handler.RegisterUser(db, newUser)
	if err != nil {
		return
	}

	fmt.Println("Registration successful.")
	fmt.Println()
}

func Login(db *sql.DB, reader *bufio.Reader) (role string, user_id int) {
	fmt.Print("\nEnter email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input proper value", err.Error())
		return
	}
	email = strings.TrimSpace(email)

	fmt.Print("Enter password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input proper value", err.Error())
		return
	}
	password = strings.TrimSpace(password)

	var loginUser = entity.User{
		Email:    email,
		Password: password,
	}

	role, user_id, err = handler.LoginUser(db, loginUser)
	if err != nil {
		return
	}

	fmt.Println("Success Login")
	fmt.Println()

	return role, user_id
}
