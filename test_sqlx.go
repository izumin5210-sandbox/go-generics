package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Connect(
		"postgres",
		fmt.Sprintf("user=%s dbname=%s sslmode=%s", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_NAME"), os.Getenv("PGSSLMODE")),
	)
	if err != nil {
		panic(err)
	}

	db.MustExec(`INSERT INTO users (name) VALUES ($1)`, "testuser")
}

func testSqlxUtils() {
	printSubSection("SqlxSelect", testSqlxSelect)
}

func testSqlxSelect() {
	snippet(
		func() interface{} {
			type User struct {
				ID   uint64
				Name string
			}
			ctx := context.Background()
			users, err := sqlxSelect[*User](ctx, db, "SELECT * FROM users")
			if err != nil {
				panic(err)
			}
			return users
		},
	)
}
