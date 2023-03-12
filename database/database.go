package database

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var err error

// Definimos una estructura CustomLogger que contiene nuestra implementación personalizada de LogWriter
type CustomLogger struct{}

func (logger CustomLogger) Print(v ...interface{}) {
	if len(v) > 1 {
		level := v[0]
		if level == "sql" {
			// Customize the output of the SQL queries
			sql := v[3].(string)
			formattedSQL := strings.Replace(sql, "\n", " ", -1)
			formattedSQL = strings.Replace(formattedSQL, "\t", "", -1)
			msg := fmt.Sprintf("\033[35m[SQL]\033[0m \033[33m%v\033[0m\n", formattedSQL)
			fmt.Print(msg)
			return
		}
	}
	// If the log level is not "sql", print the message as-is
	fmt.Println(v...)
}

func InitConnection() {

	dsn := "root:@tcp(127.0.0.1:3306)/biblioteca?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}

}

func GetConnection() *gorm.DB {
	return db
}
