package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test() int {
	return 99
}
func InitDB() (*gorm.DB, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/training?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return DB, nil
}
