package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type dncpResponseOauth struct {
	tokenType   string `json:"token_type"`
	accessToken string `json:"access_token"`
}

func main() {
	url := "https://contrataciones.gov.py/datos/api/oauth/token"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Basic MTkwMzBjNjUtNWUzZC00MDFmLWEyMmQtM2Q3OTY1YjdkOTA3OmRjOTAwYTg0LWIxZjktNDAwYi05YmVmLTk1YTA3ODkwZDk0MQ==")
	req.Header.Set("Content-Type", "application/json")
	//	fmt.Println("request Headers:", req.Header)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	//	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	var s = new(dncpResponseOauth)

	err = json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops json", err)
	}
	fmt.Println("json", s)
}
