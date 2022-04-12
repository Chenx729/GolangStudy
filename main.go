package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	db, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test_db")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = db
}

func main() {
	res, err := Db.Exec("delete from person where user_id =?", 2)
	if err != nil {
		fmt.Println("delete from person failed,", err)
		return
	}
	defer Db.Close() // 注意这行代码要写在上面err判断的下面

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("row failed", err)
	}

	fmt.Println("delete from person succeeded:", row)
}
