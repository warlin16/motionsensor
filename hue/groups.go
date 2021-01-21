package hue

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
	defer resp.Body.Close()
	fmt.Println("response status:", resp.Status)
	fmt.Println("response headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response body:", string(body))
}
