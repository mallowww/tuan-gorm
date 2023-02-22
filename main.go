package main

import (
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func main() {
	// dsn
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]?parseTime=true
	dsn := os.Getenv("DSN")
	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial)
	if err != nil {
		panic(err)
	}
	// migrator - manage database
	db.Migrator().CreateTable()
}

type Gender struct {
	ID uint
	Name string
}