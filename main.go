package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var name = "Tom"

type Response struct {
	QUOTE  string
	AUTHOR string
}

func main() {
	//greeting, err := greetings.Hello(name)
	//author, quote := greetings.GetRandomQuote()
	//if err != nil {
	//	panic(err)
	//}
	//if author == "" && quote == "" {
	//	panic("Error getting quote")
	//}
	http.HandleFunc("/random-quote", getRandomQuote)
	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		errors.New("Error starting server")
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "192.168.10.143:8080")
}

func getRandomQuote(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.api-ninjas.com/v1/quotes?category=happiness", nil)
	if err != nil {
		message := "Error creating HTTP request:"
		fmt.Println(message, err)
		return
	}

	// add API NINJAS api key to the request header
	req.Header.Add("X-Api-Key", "aKHwf19lEJwevqc/U6SaTg==XWXwjUAkrl53rfyB")
	req.Header.Add("Accept", "application/json")

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		message := "Error sending HTTP request:"
		fmt.Println(message, err)
		return
	}
	var data []Response
	err1 := json.NewDecoder(resp.Body).Decode(&data)
	if err1 != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data[0].QUOTE)
}

// Path: tom/golang-test/main_test.go
