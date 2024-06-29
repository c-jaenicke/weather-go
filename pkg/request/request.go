package request

import (
	"io"
	"log"
	"net/http"
	"os"
)

func Request(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error making request: " + err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Response body was empty.")
		os.Exit(1)
	}

	if response.StatusCode != 200 {
		log.Println("Response was code was not 200 got: " + response.Status)
		log.Println(string(responseData))
		os.Exit(1)
	}

	return responseData
}
