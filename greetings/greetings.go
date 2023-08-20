package greetings

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(GetRandomQuote())
}

type Response struct {
	QUOTE  string
	AUTHOR string
}

func GetRandomQuote() string {
	// Create an HTTP client
	client := &http.Client{}

	// Create an HTTP request authorization header containing the API key
	req, err := http.NewRequest("GET", "https://api.api-ninjas.com/v1/quotes?category=happiness", nil)
	if err != nil {
		message := "Error creating HTTP request:"
		fmt.Println(message, err)
		return message
	}

	// add API NINJAS api key to the request header
	req.Header.Add("X-Api-Key", "aKHwf19lEJwevqc/U6SaTg==XWXwjUAkrl53rfyB")
	req.Header.Add("Accept", "application/json")

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		message := "Error sending HTTP request:"
		fmt.Println(message, err)
		return message
	}
	var data []Response
	err1 := json.NewDecoder(resp.Body).Decode(&data)
	if err1 != nil {
		return "Error decoding JSON" + err1.Error()
	}
	return fmt.Sprintf("%v,%v", data[0].AUTHOR, data[0].QUOTE)
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// Return an error if name is empty.
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
