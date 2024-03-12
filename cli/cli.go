package cli

import (
	"Sport-PairProject/config"
	"fmt"
)

func RunApplication() {

	db, err := config.GetConnection()
	if err != nil {
		fmt.Println(err.Error())
		return 
	}

	fmt.Println("DB CONNECTED", db.Ping())
}