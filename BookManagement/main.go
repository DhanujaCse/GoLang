package main

import (
	"hello/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	controllers.ConnectDatabase()

	// Routes
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.POST("/employee", controllers.CreateEmployee)
	r.GET("/employee", controllers.GetAllEmployees)
	r.GET("/employee/:id", controllers.GetEmployeeById)
	r.GET("/employee/name/:employeename", controllers.GetEmployeeByEmployeeName)
	r.DELETE("/employee/:id", controllers.DeleteEmployee)
	r.PATCH("/employee/:id", controllers.UpdateEmployee)

	// Run the server
	r.Run()
}
