package connection

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	var err error

	mysqlURL := fmt.Sprintf("root:root@tcp(%s:%s)/TODOLIST?parseTime=true", viper.GetString("MYSQL_HOST"), viper.GetString("MYSQL_PORT"))

	db, err = gorm.Open("mysql", mysqlURL)

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func DB() *gorm.DB {
	return db
}

func Close() {
	db.Close()
}
