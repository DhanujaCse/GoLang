package controllers

import (
	"hello/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func PostUser(c *gin.Context) {
// 	var book models.Book
// 	c.Bind(&book)

// 	if book.Author != "" && book.Title != "" {
// 		if insert, _ := dbmap.Exec(`INSERT INTO book(Author,Title)VALUES(?,?)`, book.Author, book.Title); insert != nil {
// 			book_id, err := insert.LastInsertId()

//				if err == nil {
//					content := &models.Book{
//						Id:     book_id,
//						Author: book.Author,
//						Title:  book.Title,
//					}
//					c.JSON(201, content)
//				} else {
//					checkErr(err, "data not inserted")
//				}
//			}
//		} else {
//			c.JSON(400, gin.H{"err": "Fields are empty"})
//		}
//	}
// func PostBook(c *gin.Context) {
// 	var user models.Book
// 	c.Bind(&user)

// 	log.Println(user)

// 	if user.Author != "" && user.Title != "" {

// 		if insert, _ := dbmap.Exec(`INSERT INTO book (Author, Title) VALUES (?, ?)`, user.Author, user.Title); insert != nil {
// 			user_id, err := insert.LastInsertId()
// 			if err == nil {
// 				content := &models.Book{
// 					Id:     user_id,
// 					Author: user.Author,
// 					Title:  user.Title,
// 				}
// 				c.JSON(201, content)
// 			} else {
// 				checkErr(err, "Insert failed")
// 			}
// 		}

// 	} else {
// 		c.JSON(400, gin.H{"error": "Fields are empty"})
// 	}

// }
// func GetBookDetail(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var book models.Book
// 	err := dbmap.SelectOne(&book, "SELECT * FROM book WHERE id=? LIMIT 1", id)

// 	if err == nil {
// 		user_id, _ := strconv.ParseInt(id, 0, 64)

// 		content := &models.Book{
// 			Id:     user_id,
// 			Author: book.Author,
// 			Title:  book.Title,
// 		}
// 		c.JSON(200, content)
// 	} else {
// 		c.JSON(404, gin.H{"error": "book not found"})
// 	}
// }

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// GET /books
// Find all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// GET /books/:id
// Find a book
func FindBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// POST /books
// Create new book
func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

type CreateEmployeeInput struct {
	Employeename        string `json:"employeename"`
	Employeedesignation string `json:"employeedesignation"`
}

func CreateEmployee(c *gin.Context) {
	var input CreateEmployeeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	employee := models.Employee{EmployeeName: input.Employeename, EmployeeDesignation: input.Employeedesignation}

	DB.Create(&employee)
	c.JSON(http.StatusOK, gin.H{"data": employee})

}
func GetAllEmployees(c *gin.Context) {
	var employees []models.Employee
	DB.Find(&employees)
	c.JSON(http.StatusOK, gin.H{"data": employees})
}
func GetEmployeeById(c *gin.Context) {
	var employee models.Employee

	if err := DB.Where("id=?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "data not found"})
		return
	}

	c.JSON(http.StatusFound, gin.H{"data": employee})
}
func GetEmployeeByEmployeeName(c *gin.Context) {
	var employee []models.Employee
	if err := DB.Where("employee_name=?", c.Param("employeename")).Find(&employee).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"err": "no data found"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

func DeleteEmployee(c *gin.Context) {
	var employee models.Employee
	if err := DB.Where("id=?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"err": "no data found"})
		return
	}
	DB.Delete(&employee)
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

func UpdateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := DB.Where("id=?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"err": "data not found"})
		return

	}
	var input CreateEmployeeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	DB.Model(&employee).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": employee})

}
