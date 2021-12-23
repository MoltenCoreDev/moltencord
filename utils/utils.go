package utils

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

var (
	auth string
)

func SetToken(token string) {
	auth = token
}

func MakeRequest(method string, url string, data []byte) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest(string(method), url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", auth)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if resp.StatusCode > 199 && resp.StatusCode < 299 {
		return resp, err
	} else {
		msg, _ := io.ReadAll(resp.Body)
		return nil, errors.New("Error " + resp.Status + ":" + string(msg))
	}

}
