package common

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is a global.for controller to use
var DB *gorm.DB

// InitDB is a function of init database
func InitDB() {

	dbuser := viper.GetString(`database.dbuser`)
	dbpass := viper.GetString(`database.dbpass`)
	dbhost := viper.GetString(`database.dbhost`)
	dbport := viper.GetString(`database.dbport`)
	dbname := viper.GetString(`database.dbname`)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbport, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Faild to connect database")
	}
	DB = db

}

// GetDB for controller to use
func GetDB() *gorm.DB {
	return DB
}
