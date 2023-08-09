package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func MakeRequest(endpoint string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)

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

func BuildFilterURL(secure bool, endpoint string, port int, mode string, filter []string) string {
	protocol := "http"

	if secure {
		protocol = "https"
	}

	readyFilter := ""

	if filter[0] != "" {
		for i, v := range filter {
			escapedFilter := url.QueryEscape(v)

			if i == 0 {
				readyFilter = "?filter=" + escapedFilter
			} else {
				readyFilter += "&filter=" + escapedFilter
			}

		}

	}

	return fmt.Sprintf("%s://%s:%d/api/v2/%s%s", protocol, endpoint, port, mode, readyFilter)
}

func BuildNamedURL(secure bool, endpoint string, port int, mode string, name string) string {
	protocol := "http"

	if secure {
		protocol = "https"
	}

	return fmt.Sprintf("%s://%s:%d/api/v2/%s?name=%s", protocol, endpoint, port, mode, name)
}
