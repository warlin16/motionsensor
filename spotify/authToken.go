package spotify

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"motionsensor/config"
	"net/http"
	"strings"
)

const method = "POST"

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

// SetAuthToken sets the auth token required to talk to spotify api
func SetAuthToken(c *config.Config) {
	client := &http.Client{}
	encodedCreds := b64.StdEncoding.EncodeToString([]byte(c.SpoClientID + ":" + c.SpoClientSecret))
	tr := tokenResponse{}

	payload := strings.NewReader("grant_type=client_credentials")

	req, err := http.NewRequest(method, c.SpotifyTokenURL, payload)

	if err != nil {
		log.Fatalf("Failed to create POST request to token url: " + err.Error())
	}

	req.Header.Add("Authorization", "Basic "+encodedCreds)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)

	if err != nil {
		log.Fatalf("Failed to send POST request to token url: " + err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Failed to read POST request: " + err.Error())
	}

	json.Unmarshal(body, &tr)
	c.SpoOauthToken = tr.AccessToken
	fmt.Println("this is the config struct now:", c.SpoOauthToken)
}
