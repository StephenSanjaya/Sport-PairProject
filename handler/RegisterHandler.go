package handler

import (
	"Sport-PairProject/entity"
	"database/sql"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *sql.DB, newUser entity.User) error {
	// Check email format
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	if !emailRegex.MatchString(newUser.Email) {
		return fmt.Errorf("invalid email format")
	}

	// Hash user password with GenerateFormPassword method
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Insert user data into the database
	_, err = db.Exec("INSERT INTO Users(Username, Email, Password, Address) VALUES (?, ?, ?, ?)", newUser.Username, newUser.Email, hashedPassword, newUser.Address)
	if err != nil {
		return err
	}

	return nil
}
