package router

import (

	"fmt"
	"go_book/book"
	"go_book/transaction"
	"go_book/externalApi"
	"github.com/gofiber/fiber/v2"
)

/* {"docs":[{"_id":"5cf5805fb53e011a64671582","name":"The Fellowship Of The Ring"},{"_id":"5cf58077b53e011a64671583","name":"The Two Towers"},{"_id":"5cf58080b53e011a64671584","name":"The Return Of The King"}],"total":3,"limit":1000,"offset":0,"page":1,"pages":1}*/


func PostWorld(c *fiber.Ctx) error {
	// curl http://localhost:5000/hello/lijun
	name := c.Params("name")
	// Alternatively you can get POST data from 'queries'
	// e.g.
	// c.Query("name")
	// curl localhost:3000/hello?name="world"
	fmt.Println(name)
	// if name != "World" {
	// 	c.JSON("Error")
	// }

	if name == "" {
		c.SendStatus(400)
		return c.JSON("empty")
	}
	return c.JSON("Hello, " + name)
}

var name = "World"

func SetupRoutes(app *fiber.App) {
	// Hello World Routes
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON("Hello, " + name)
	})

	// curl -X POST http://localhost:5000/hello/lijun
	app.Post("/hello/:name?", PostWorld)

	//Transaction Routes
	app.Get("/api/transaction", transaction.GetTransactions)
	app.Post("/api/transaction", transaction.NewTransaction)

	// Book Routes
	app.Get("/api/book", book.GetBooks)
	app.Get("/api/book/:id", book.GetBook)
	app.Post("/api/book", book.NewBook)
	app.Delete("/api/book/:id", book.DeleteBook)

	// 	Endpoint GET /External-API
	app.Get("/api/external-api", externalApi.ExternalAPI)
	app.Get("/api/graphQLPost", func(c *fiber.Ctx) error {

		externalApi.GQLcreateUsers()
		return c.JSON("Post")

	})
	app.Get("/api/graphQLGet", func(c *fiber.Ctx) error {

		externalApi.GQLUsers()
	
		return c.JSON("get")

	})
	
	

}