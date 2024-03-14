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
	defer db.Close()
}
