package handler

import (
	"Sport-PairProject/entity"
	"database/sql"
)

func LoginUser(db *sql.DB, newUser entity.User) (role string, user_id int) {
	//login logic

	return role, user_id
}
