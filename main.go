package main

import "fmt"
import "os"
import "flag"
import "net/http"
import "encoding/json"

type Location struct {
	Lat, Long float32
	Accuracy int
}

func main()  {
	var googleApiKey string
	
	apiKeyEnv := os.Getenv("GOOGLE_API_KEY")
	if apiKeyEnv != "" {
		googleApiKey = apiKeyEnv
	}

	apiKeyPtr := flag.String("apikey", "", "Google API Key")
	flag.Parse()
	if *apiKeyPtr != "" {
		googleApiKey = *apiKeyPtr
	}

	if googleApiKey == "" {
		flag.Usage()
		os.Exit(1)
	}

	apiRequest := fmt.Sprintf("https://www.googleapis.com/geolocation/v1/geolocate?key=%s", googleApiKey)
	resp, err := http.Post(apiRequest, "text/plain", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var jsonResponse map
	json.NewDecoder(resp.Body).Decode(jsonResponse)
}