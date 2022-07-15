package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key := os.Getenv("API_KEY")
	secret := os.Getenv("API_SECRET")
	client := smtpd.New(key, secret)

	p, err := client.GetAllProfiles()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(PrettyJSON(p))
}

func PrettyJSON(input interface{}) string {
	b, err := json.Marshal(input)
	if err != nil {
		log.Fatal(err)
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, b, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	return prettyJSON.String()
}
