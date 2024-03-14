package handler

import (
	"Sport-PairProject/entity"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func LoginUser(db *sql.DB, newUser entity.User) (role string, user_id int, err error) {
	var email, password string
	err = db.QueryRow("SELECT UserID, Role, Email, Password FROM Users WHERE Email = $1", newUser.Email).Scan(&user_id, &role, &email, &password)
	if err != nil {
		fmt.Println("User not found: ", err)
		return "", -1, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(newUser.Password))
	if err != nil { // for compare input password with hashed password in database with CompareHashAndPassword method
		return "", -1, err
	}

	return role, user_id, nil
}
