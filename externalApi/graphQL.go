package externalApi

import (
	"fmt"
	"log"
	"net/http"
	// "github.com/gofiber/fiber/v2"
	"strings"
	"io/ioutil"
	// "time"
	// "bytes"
	// "encoding/json"
)

/*

curl \
  -H "Content-Type: application/json" \
  -d '{ "query": "mutation { createUser(input: { name: \"Odour\"  id: \"1\"}) { id name } }" }' \
  http://localhost:8080/query

*/

func GQLcreateUsers() string {
	
	idq := 1 
	// nameq := "jun"

	str := fmt.Sprintf(`{"query": "mutation{ createUser ( input: { id: %d, name: "a"}) {id  name} }"}` , idq )
	
	fmt.Println(str)
	body := strings.NewReader(str)
	fmt.Println(body)
	req, err := http.NewRequest("POST", "http://localhost:8080/query", body)
	
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	
	// defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	sb := string(data)
	log.Println(sb)

	log.Println(req)

	
	return "not woking"


}

func GQLUsers() string {
	str := `{ "query": "users{ name }" }`
	
	body := strings.NewReader(str)
	fmt.Println(str)
	req, err := http.NewRequest("GET", "http://localhost:8080/query", body)
	
	if err != nil {
		log.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")
	
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	
	defer resp.Body.Close()
	log.Println(resp)
	data, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return c.SendStatus(500)
	// }
	// defer resp.Body.Close()
	sb := string(data)
	// log.Printf(sb)
	log.Println(err)
	log.Println(sb)
	return "not woking"

	// jsonData := map[string]string{
    //     "query": `
    //         { 
    //             users {
    //                 name
    //                 id
    //             }
    //         }
    //     `,
    // }
    // jsonValue, _ := json.Marshal(jsonData)
    // request, err := http.NewRequest("POST", "http://localhost:8080/query", bytes.NewBuffer(jsonValue))
	// if err != nil {
	// 	log.Println(err)
	// }
    // client := &http.Client{Timeout: time.Second * 10}
    // response, err := client.Do(request)

    // if err != nil {
	// 	fmt.Printf("The HTTP request failed with error %s\n", err)
    // }

	// defer response.Body.Close()

    // data, _ := ioutil.ReadAll(response.Body)
    // fmt.Println(string(data))

	// return c.JSON("not woking")

}

