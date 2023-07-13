package controllers

import (
	"hello/models"
	"log"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

// var dbmap = initDb()

// func initDb() *gorp.DbMap {
// 	db, err := sql.Open("mysql", "root:Dhanu123@tcp(localhost:3306)/prg")
// 	checkErr(err, "sql connection failed")

// 	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF"}}
// 	err = dbmap.CreateTablesIfNotExists()
// 	checkErr(err, "table not created")
// 	return dbmap
// }

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:Dhanu123@tcp(127.0.0.1:3306)/prg?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Book{})
	database.AutoMigrate(&models.Employee{})

	DB = database
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
