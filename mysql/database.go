package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	fmt.Println("Go MySQL connect Open")
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"), //TODO os.Getenv
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}
	var err error
	if Db, err = sql.Open("mysql", cfg.FormatDSN()); err != nil {
		log.Fatal("Db open error:", err.Error())
	}

	checkConnect(100)
	fmt.Println("Go MySQL connect Close")
}

func checkConnect(count uint) {
	var err error
	if err = Db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		fmt.Printf("retry...count:%v\n", count)
		checkConnect(count)
	}
}
