package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AuthHandler is working!")
}

// Encode keys in base64
func KeysToBase64(applicationKey, clientKey string) string {
	return base64.StdEncoding.EncodeToString([]byte(applicationKey + ":" + clientKey))
}

// Get tokens using client credentials
func GetTokens(base64 string) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", "https://api-demo.poweroffice.net/oauth/token", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Add("Authorization", "Basic "+base64)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ReadAll: ", err)
		return
	}
	json.Unmarshal(body, &result)
	log.Println("Response: ", result)
}
