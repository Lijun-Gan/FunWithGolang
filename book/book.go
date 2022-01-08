package book

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	
	"go_book/database"
)

type Book struct {
	gorm.Model
	Title string `json:"title"`
	// Title  string `json:"title" valid:"required"`
	// Title  string `json:"title" gorm:"not null"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	return c.JSON(book)
}

// func NewBook(c *fiber.Ctx) {
// 	db := database.DBConn
// 	var book Book
// 	book.Title = "1984"
// 	book.Author = "George Orwell"
// 	book.Rating = 5
// 	db.Create(&book)
// 	c.JSON(book)
// }

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		// c.Status(503).Send(err)
		c.Status(503)
		return c.JSON("Error")
	}
	db.Create(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println(id)
	db := database.DBConn

	var book Book
	db.First(&book, id)

	// check the id if exits ?
	if book.ID == 0 {

		c.Status(503)
		return c.JSON("No Book Found with ID")

	}

	fmt.Println(book.Author)
	// if book.Title == "" {
	//     c.Status(500).Send("No Book Found with ID")
	//     return
	// }

	db.Delete(&book)
	return c.JSON("Book Successfully deleted")
}
