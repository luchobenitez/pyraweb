package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type oauthPayload struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv("SECRET_KEY")

	url := "https://contrataciones.gov.py/datos/api/oauth/token"
	fmt.Println("URL: ", url)

	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Basic "+secretKey)

	res, postErr := http.DefaultClient.Do(req)

	if postErr != nil {
		log.Fatal(postErr)
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		log.Fatal(err)
	}

	//fmt.Println(res)
	//fmt.Println(string(body))

	responsePayload := oauthPayload{}

	jsonErr := json.Unmarshal(body, &responsePayload)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println("token type:", responsePayload.TokenType)
	fmt.Println("access token:", responsePayload.AccessToken)
}
