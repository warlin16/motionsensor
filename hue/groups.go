package hue

import (
	"bytes"
	"log"
	"net/http"
	"os"
)

// SetLivingRoomBrightness sets the living room brightness
func SetLivingRoomBrightness(j []byte) {
	req, _ := http.NewRequest("PUT", os.Getenv("HUE_API_URL")+"/groups/2/action", bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf(err.Error())
	}
	resp.Body.Close()
}
