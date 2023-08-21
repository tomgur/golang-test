package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Response struct {
	QUOTE  string
	AUTHOR string
}

func main() {
	http.HandleFunc("/random-quote", getRandomQuote)
	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		err1 := errors.New("error starting server")
		if err1 != nil {
			fmt.Println("error starting server")
		}
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://192.168.10.143:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
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
		fmt.Println(message, err.Error())
		return
	}
	var data []Response
	err1 := json.NewDecoder(resp.Body).Decode(&data)
	if err1 != nil {
		message := "Error creating JSON decoder:"
		fmt.Println(message, err1.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	err2 := json.NewEncoder(w).Encode(data[0].QUOTE)
	if err2 != nil {
		message := "Error encoding JSON\n"
		fmt.Println(message, err2.Error())
	}
}
