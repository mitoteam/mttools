package mttools

import (
	"io"
	"net/http"
)

// Performs GET request, returns response body as a string.
func SimpleGet(url string) (string, error) {
	res, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	bodyData, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(bodyData), nil
}
