package handler

import "database/sql"

func IncreaseUserBalanceByID(db *sql.DB, userID int, amountToAdd float64) error {
	_, err := db.Exec("UPDATE Users SET Balance = Balance + $1 WHERE UserID = $2", amountToAdd, userID)
	return err
}