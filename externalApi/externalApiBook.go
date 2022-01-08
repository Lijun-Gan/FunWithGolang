package externalApi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gofiber/fiber/v2"

)

type ExternalData struct {
	Id 		string    `json:"_id"`
	Name 	string    `json:"name"`
}

type ExternalDocs struct {
	Docs []ExternalData  `json:"docs"`
}

func ExternalAPI(c *fiber.Ctx) error {
	// https://the-one-api.dev/v2/book

	resp, err := http.Get("https://the-one-api.dev/v2/book")

	if err != nil {
		return c.SendStatus(500)
	}
	// JSON Marshalling//UnMarshalling
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.SendStatus(500)
	}
	defer resp.Body.Close()
	// sb := string(data)
	// log.Printf(sb)
	jsonData := new(ExternalDocs) 
	// Unmarshal parse raw JSON data in the form of []byte variables
	jsonErr := json.Unmarshal([]byte(data), &jsonData)
	
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	// fmt.Println(jsonData)
	// fmt.Println(data)
	return c.JSON(jsonData)
	// return c.SendStatus(500)
}