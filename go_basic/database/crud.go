package database

import (
	"database/sql"
	"fmt"
	"time"
)

func Query(db *sql.DB) {
	rows, err := db.Query("select * from user")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var username, password, email string
		var createTime time.Time
		err := rows.Scan(&id, &username, &password, &email, &createTime)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(id, username, password, email, createTime)
	}

}
