package hue

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"motionsensor/config"
	"net/http"
	"os"
)

var c config.Config

// Bridge slice of structs will contain info about hue bridges in your local network
type Bridge []struct {
	ID                string `json:"id,omitempty"`
	Internalipaddress string `json:"internalipaddress,omitempty`
}

// GetBridgeInfo fetches details about local hue bridges
func GetBridgeInfo() {
	c.FetchHueBridgeURL()
	resp, err := http.Get(c.HueBridgeURL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	bridges := Bridge{}

	err = json.Unmarshal(body, &bridges)
	os.Setenv("HUE_API_URL", "http://"+bridges[0].Internalipaddress+"/api/"+c.HueUsername)
}
