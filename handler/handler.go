package handler

import (
	"Sport-PairProject/entity"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var users = make(map[string]entity.Users)

func ShowForm() {
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
			Register(reader)
		case "2":
			Login(reader)
		case "3":
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid option. Please choose again.")
		}
	}
}

func Register(reader *bufio.Reader) {
	fmt.Println("Registration")
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	if _, exists := users[username]; exists {
		fmt.Println("Username already exists.")
		return
	}

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	newUser := entity.Users{
		Username: username,
		Password: password,
	}

	users[username] = newUser
	fmt.Println("Registration successful.")
}

func Login(reader *bufio.Reader) {
	fmt.Println("Login")
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	if storedUser, exists := users[username]; exists {
		if storedUser.Password == password {
			fmt.Println("Login successful.")
			return
		}
	}

	fmt.Println("Invalid username or password.")
}
