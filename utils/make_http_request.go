package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func MakeHTTPRequest[T any](fullURL string, httpMethod string, headers map[string]string, queryParameters url.Values, body io.Reader, response T) (T, error) {
	client := http.Client{}
	u, err := url.Parse(fullURL)
	if err != nil {
		return response, err
	}

	if httpMethod == "GET" {
		q := u.Query()

		for key, value := range queryParameters {
			q.Set(key, strings.Join(value, ","))
		}

		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(httpMethod, u.String(), body)
	if err != nil {
		return response, err
	}

	for k,v := range headers {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)

	if err != nil {
		return response, err
	}

	if res == nil {
		return response, fmt.Errorf("HTTP response is nil")
	}

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return response, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return response, fmt.Errorf("HTTP request failed with status code: %d, responseData: %s on url : %s", res.StatusCode, responseData, u.String())
	}

	var responseObject T
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return response, err
	}

	return responseObject, nil
}