package auth

import (
	"log"
	"os"

	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type TokenData struct {
	Audience     string `json:"audience"`
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

var GetToken = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	client := new(http.Client)
	url := "https://" + os.Getenv("AUTH0_DOMAIN") + "/oauth/token"
	data, err := json.Marshal(TokenData{Audience: os.Getenv("AUTH0_API_IDENTIFIER"), GrantType: "client_credentials", ClientId: os.Getenv("AUTH0_CLIENT_ID"), ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET")})
	if err != nil {
		log.Fatal(err)
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var tokenResponse TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		log.Fatal(err)
	}
	jsonResponse, err := json.Marshal(tokenResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
})
