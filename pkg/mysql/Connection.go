package mysql

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	RW *gorm.DB
)

func NewDB() (*gorm.DB, error) {
	DBMS := "mysql"
	mySqlConfig := &mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "dbname",
		AllowNativePasswords: true,
		Params: map[string]string{
			"parseTime": "true",
		},
	}

	return gorm.Open(DBMS, mySqlConfig.FormatDSN())
}

func init() {
	readWriteConnection, err := NewDB()
	if err != nil {
		log.Panicf("데이터베이스 연결에 실패했습니다: %v", err)
	}
	RW = readWriteConnection
}
