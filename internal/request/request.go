package request

import (
	"fmt"
	"io"
	"net/http"
)

// Request makes an HTTP GET request and returns the response body as a byte slice.
func Request(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("request to %s returned status code %d: %s", url, response.StatusCode, string(responseData))
	}

	return responseData, nil
}
