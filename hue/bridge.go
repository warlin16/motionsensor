package hue

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Bridge slice of structs will contain info about hue bridges in your local network
type Bridge []struct {
	ID                string `json:"id,omitempty"`
	Internalipaddress string `json:"internalipaddress,omitempty`
}

// GetBridgeInfo fetches details about local hue bridges
func GetBridgeInfo() Bridge {
	resp, err := http.Get("someurl")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	bridges := Bridge{}

	err = json.Unmarshal(body, &bridges)
	return bridges
}
