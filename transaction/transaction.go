package transaction

import (
	"github.com/gofiber/fiber/v2"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go_book/database"
)

// type Transaction struct {
// 	gorm.Model
// 	AccountID 	int `json:"accountID"`
// 	Balance 	int `json:"balance"`
// }

type Transaction struct {
	gorm.Model
	TransactionID int `json:"transactionID"`

	Data struct {
		AccountID int    `json:"accountID"`
		Balance   int    `json:"balance"`
		Date      string `json:"dateUpdated"`
	} `json:"Data"`
}

// type data struct {
// 	AccountID int    `json:"accountID"`
// 	Balance   int    `json:"balance"`
// 	Date      string `json:"dateUpdated"`
// }

// TransactionID and DateUpdated => these will be auto generated
// accountID -- is this from another table ?  is the DateUpdated means when the transaction happened?

func GetTransactions(c *fiber.Ctx) error {
	db := database.DBConn
	var transactions []Transaction

	db.Find(&transactions)
	return c.JSON(transactions)
}

func NewTransaction(c *fiber.Ctx) error {
	db := database.DBConn
	var transaction Transaction
	transaction.TransactionID = 29099
	transaction.Data.AccountID = 111
	transaction.Data.Balance = 123
	db.Create(&transaction)
	return c.JSON(transaction)
}

func DeleteTransaction(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var transaction Transaction
	db.First(&transaction, id)

	// if transaction.AccountID == "" {
	//     c.Status(500).Send("No Book Found with ID")
	//     return
	// }
	db.Delete(&transaction)
	return c.JSON("Transaction Successfully deleted")

}
