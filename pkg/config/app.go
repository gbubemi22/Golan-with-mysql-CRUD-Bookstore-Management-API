package config

import(
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	//  _ "github.com/jinzhu/gorm/dialects/postgres"
	//  _ "github.com/jinzhu/gorm/dialects/sqlite"
	//"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var (
	db * gorm.DB
)

func Connect() {
     d, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
	                              
	if err != nil {
		panic(err)

	}
	db = d
     
}

func GetDB() *gorm.DB {
     return db
} 