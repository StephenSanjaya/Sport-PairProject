package handler

import (
	"Sport-PairProject/entity"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func LoginUser(db *sql.DB, newUser entity.User) (role string, user_id int) {
	var email, password string
	err := db.QueryRow("SELECT UserID, Role, Email, Password FORM Users WHERE Email = ?", newUser.Email).Scan(&user_id, &role, &email, &password)
	if err != nil {
		fmt.Println("User not found: ", err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(password))
	if err != nil { // for compare input password with hashed password in database with CompareHashAndPassword method
		return
	}

	return role, user_id
}
