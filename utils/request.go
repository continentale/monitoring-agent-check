package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func MakeRequest(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, errors.New("Unable to initialize new request: " + err.Error())
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, errors.New("Unable to make request: " + err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, errors.New("Unable to read body from request: " + err.Error())
	}

	return body, nil
}

func BuildURL(secure bool, url string, port int, mode string, filter string) string {
	protocol := "http"

	if secure {
		protocol = "https"
	}

	if filter != "" {
		filter = "?filter=" + filter
	}

	return fmt.Sprintf("%s://%s:%d/api/v2/%s%s", protocol, url, port, mode, filter)
}
