package handler

import (
	"Sport-PairProject/entity"
	"context"
	"database/sql"
	"fmt"
	"time"
)

func GetUserReport(db *sql.DB) (user_report []entity.UserReport, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var ur entity.UserReport

	query := "SELECT UserID, Username, Email, Address, Balance FROM Users"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		fmt.Println(err.Error())
		return []entity.UserReport{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ur.UserID, &ur.Username, &ur.Email, &ur.Address, &ur.Balance)
		if err != nil {
			fmt.Println(err.Error())
			return []entity.UserReport{}, err
		}
		user_report = append(user_report, ur)
	}

	return user_report, nil
}
