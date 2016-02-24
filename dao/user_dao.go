package dao

import ()

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllUser() ([]string, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/registry")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	rows, err := db.Query("select username from user")
	if err != nil {
		return nil, err
	}

	var users []string

	for rows.Next() {
		var user string
		if err := rows.Scan(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
