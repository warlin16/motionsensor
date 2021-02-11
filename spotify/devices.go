package spotify

import (
	"fmt"
	"io/ioutil"
	"log"
	"motionsensor/config"
	"net/http"
)

// GetDevices gets local devices
func GetDevices(c *config.Config) {
	println("This will get the local devices....")
	req, err := http.NewRequest("GET", c.SpotifyURL+"me/player/devices", nil)
	if err != nil {
		log.Fatalf("error constructing http req: " + err.Error())
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.SpoOauthToken)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to fetch local spotify devices: " + err.Error())
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("The response of the body:", string(body))
}
