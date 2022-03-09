package connection

import (
	"demo/app/domain"
	"demo/infra/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func ConnectDb() {
	dbConfig := config.Db()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.Schema)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connection is successful")

	database = db
	err = database.AutoMigrate(
		&domain.User{},
	)
	if err != nil {
		panic(err)
	}
}

func Db() *gorm.DB {
	return database
}
