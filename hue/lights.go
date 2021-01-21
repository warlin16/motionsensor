package hue

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// GetLights fetches list of active lights for a specific bridge
func GetLights() {
	fmt.Println(os.Getenv("HUE_API_URL"))
	resp, err := http.Get(os.Getenv("HUE_API_URL") + "/lights")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println(resp.Body)
}
