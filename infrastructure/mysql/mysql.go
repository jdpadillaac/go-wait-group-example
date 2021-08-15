package appmysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jdpadillaac/go-waitgroups-example/internal"
	"log"
)

var (
	db *sql.DB
)

func SetCnn(c *internal.AppConfig) {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true", c.MySqlUser, c.MySqlPassword,
		c.MySqlServer, c.MySqlPort, c.MySqlDataBase)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func GetCnn() *sql.DB {
	return db
}
