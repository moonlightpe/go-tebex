package main

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

const TebexURL = "https://plugin.tebex.io/"

// New starts a new instance to the Tebex API
func New(secret string) (Session, error) {
	if secret == "" {
		return Session{}, errors.New("your secret can't be nothing")
	}
	return Session{Secret: secret}, nil
}

func (s Session) dial(endpoint string, request string, data []byte) (body []byte, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", TebexURL+endpoint, nil)

	switch {
	case request == "post":
		req, _ = http.NewRequest("POST", TebexURL+endpoint, bytes.NewBuffer(data))
	case request == "patch":
		req, _ = http.NewRequest("PATCH", TebexURL+endpoint, bytes.NewBuffer(data))
	case request == "delete":
		req, _ = http.NewRequest("DELETE", TebexURL+endpoint, nil)
	default:
	}

	req.Header.Add("X-Tebex-Secret", s.Secret)
	req.Header.Add("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 202 && resp.StatusCode != 204 {
		body, _ = io.ReadAll(resp.Body)
		err = errors.New(string(body))
		return
	}

	if resp.Body != nil {
		body, _ = io.ReadAll(resp.Body)
		return body, nil
	}

	defer resp.Body.Close()
	return nil, nil
}
