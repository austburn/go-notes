package main

import "fmt"
import "os"
import "flag"

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

	fmt.Printf("Here is my api key %s\n", googleApiKey)
}