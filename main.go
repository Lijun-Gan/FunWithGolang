package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go_book/router"
	"go_book/book"
	"go_book/database"
	"go_book/transaction"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// "encoding/json"
	
)

// func serveStatic(app *fiber.App) {
// 	app.Static("/", "./build")
// }

// why cannot I move this to database?

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&book.Book{})
	database.DBConn.AutoMigrate(&transaction.Transaction{})
	fmt.Println("Database Migrated")
}

func main() {
	//server
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	//Handle Cors
	app.Use(cors.New())
	//Serve the build file
	// serveStatic(app)

	router.SetupRoutes(app)	//Setup Routes

	//Heroku automatically assigns a port our web server. If it   //fails we instruct it to use port 5000
	port := os.Getenv("PORT")
	if port == "" {
		port = ":5000"
	}

	log.Printf("Listening on port %s\n\n", port)
	log.Fatal(app.Listen(port))
}