package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math"
	"net/http"
	"os"
)

type quoteResponse struct {
	QUOTE    string
	AUTHOR   string
	CATEGORY string
}
type btcResponse struct {
	Time         string  `json:"time"`
	AssetIDBase  string  `json:"asset_id_base"`
	AssetIDQuote string  `json:"asset_id_quote"`
	Rate         float64 `json:"rate"`
}

func main() {
	port := ":8080"

	http.HandleFunc("/random-quote", getRandomQuote)
	http.HandleFunc("/bitcoin-price", getBitcoinPrice)
	http.HandleFunc("/api/recipe/search", searchRecipe)
	//http.HandleFunc("/mongo/ping", cacheQuote)
	fmt.Printf("Server is running on port %d...\n", port)
	certFile := "chain.pem"
	keyFile := "key.pem"

	// Start the HTTPS server
	err := http.ListenAndServeTLS(port, certFile, keyFile, nil)
	if err != nil {
		panic(err)
	}
}

func pingMongo(w http.ResponseWriter, r *http.Request) {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://tomgur80:FwuOJN5Y2Rw05X5j@cluster0.gc8xx0j.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func searchRecipe(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	//client := &http.Client{}
	//
	//return "success"
}
func cacheQuote(quote quoteResponse) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	var uri string
	if uri = os.Getenv("MONGODB_URI"); uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// begin insertOne
	coll := client.Database("mui").Collection("quotes")

	result, err := coll.InsertOne(context.TODO(), quote)
	if err != nil {
		panic(err)
	}
	// end insertOne
	fmt.Println("Inserted a single document: ", result.InsertedID)
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
	var data []quoteResponse
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
	for i := 0; i < len(data); i++ {
		cacheQuote(data[i])
	}
}

func getBitcoinPrice(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://rest.coinapi.io/v1/exchangerate/BTC/USD", nil)
	if err != nil {
		message := "Error creating HTTP request:"
		fmt.Println(message, err)
		return
	}

	req.Header.Add("X-CoinAPI-Key", "709C6CD6-EF67-4261-A974-BD5E3BDEE52F")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		message := "Error sending HTTP request:"
		fmt.Println(message, err.Error())
		return
	}
	defer resp.Body.Close()

	var responseData btcResponse

	decoder := json.NewDecoder(resp.Body)
	decodeErr := decoder.Decode(&responseData)
	if decodeErr != nil {
		message := "Error decoding JSON:"
		fmt.Println(message, decodeErr.Error())
		return
	}

	fmt.Println("Time: ", responseData.Time)
	fmt.Println("Asset ID Base: ", responseData.AssetIDBase)
	fmt.Println("Asset ID Quote: ", responseData.AssetIDQuote)
	fmt.Println("Rate: ", responseData.Rate)

	w.Header().Set("Content-Type", "application/json")
	encoderError := json.NewEncoder(w).Encode(toFixed(responseData.Rate, 2))
	if encoderError != nil {
		message := "Error encoding JSON\n"
		fmt.Println(message, encoderError.Error())
	}

}
