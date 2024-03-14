package handler

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func runDB() (db *sql.DB, err error) {
	psqlInfo := "host=ec2-18-211-172-50.compute-1.amazonaws.com port=5432 user=uapbpjeueionou password=02edc924245171009c91a45a08f5323f4cf63cd3bb489eed852ff0f435bbd9f6 dbname=d8ij6dggri9pi5 sslmode=require"

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func TestGetCurrentStock(t *testing.T) {
	db, err := runDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	datas := []struct {
		current_stock  int
		expected_stock int
	}{
		{
			GetCurrentStock(db, 1),
			18,
		},
	}

	for _, v := range datas {
		if v.expected_stock != v.current_stock {
			t.Errorf("Current stock is %d while the expected is %d", v.current_stock, v.expected_stock)
		}
	}
}
