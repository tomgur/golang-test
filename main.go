package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	QUOTE    string
	AUTHOR   string
	CATEGORY string
}

type BitcoinResponse struct {
	PRICE string
}

func main() {
	port := 8080
	http.HandleFunc("/random-quote", getRandomQuote)
	fmt.Printf("Server is running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("error starting server")
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func getRandomQuote(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.api-ninjas.com/v1/quotes", nil)
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
		message := "Error creating JSON decoder"
		fmt.Println(message, err.Error())
	}
	if len(data) > 0 {
		for i := 0; i < len(data); i++ {
			fmt.Println("quote: ", data[i].QUOTE)
			fmt.Println("author: ", data[i].AUTHOR)
			fmt.Println("category: ", data[i].CATEGORY)
		}
	} else {
		message := "No data found"
		fmt.Println(message)
	}
	w.Header().Set("Content-Type", "application/json")
	err2 := json.NewEncoder(w).Encode(data[0])
	if err2 != nil {
		message := "Error encoding JSON\n"
		fmt.Println(message, err2.Error())
	}
}

func getBitcoinPrice(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http", nil)
	if err != nil {
		message := "Error creating HTTP request:"
		fmt.Println(message, err)
		return
	}

	// add API NINJAS api key to the request header
	req.Header.Add("X-CMC_PRO_API_KEY", "df32b477-7561-40df-acd0-e1bf88b709d1")
	req.Header.Add("Accept", "application/json")

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		message := "Error sending HTTP request:"
		fmt.Println(message, err.Error())
		return
	}
	var bitcoinData []BitcoinResponse
	bitcoinError := json.NewDecoder(resp.Body).Decode(&bitcoinData)
	if bitcoinError != nil {
		message := "Error creating JSON decoder"
		fmt.Println(message, err.Error())
	}
	if len(bitcoinData) > 0 {
		for i := 0; i < len(bitcoinData); i++ {
			fmt.Println("Price: ", bitcoinData[i].PRICE)
		}
	} else {
		message := "No data found"
		fmt.Println(message)
	}
	w.Header().Set("Content-Type", "application/json")
	encoderError := json.NewEncoder(w).Encode(bitcoinData[0])
	if encoderError != nil {
		message := "Error encoding JSON\n"
		fmt.Println(message, encoderError.Error())
	}
}

func getIlsPrice(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		message := "Error creating HTTP request:"
		fmt.Println(message, err)
		return
	}

	// add API NINJAS api key to the request header
	req.Header.Add("X-CMC_PRO_API_KEY", "df32b477-7561-40df-acd0-e1bf88b709d1")
	req.Header.Add("Accept", "application/json")

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		message := "Error sending HTTP request:"
		fmt.Println(message, err.Error())
		return
	}
	var bitcoinData []BitcoinResponse
	bitcoinError := json.NewDecoder(resp.Body).Decode(&bitcoinData)
	if bitcoinError != nil {
		message := "Error creating JSON decoder"
		fmt.Println(message, err.Error())
	}
	if len(bitcoinData) > 0 {
		for i := 0; i < len(bitcoinData); i++ {
			fmt.Println("Price: ", bitcoinData[i].PRICE)
		}
	} else {
		message := "No data found"
		fmt.Println(message)
	}
	w.Header().Set("Content-Type", "application/json")
	encoderError := json.NewEncoder(w).Encode(bitcoinData[0])
	if encoderError != nil {
		message := "Error encoding JSON\n"
		fmt.Println(message, encoderError.Error())
	}
}
