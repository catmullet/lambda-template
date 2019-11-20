package client

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

// Post Makes a post call url and returns response and error
func Post(username, password, url string, body []byte, headers map[string]string) (*http.Response, error) {
	return do(username, password, "POST", url, bytes.NewBuffer(body), nil, headers)
}

// Get Makes a get call to url and returns response and error
func Get(username, password, url string, params map[string]string) (*http.Response, error) {
	return do(username, password, "GET", url, nil, params, nil)
}

// Makes the request based on the http method passed and values.
func do(username, password, method, url string, body io.Reader, params map[string]string, headers map[string]string) (*http.Response, error) {

	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	setParams(req, params)
	setRequestHeaders(req, headers)
	setBasicAuth(req, username, password)

	return getHttpClient().Do(req)
}

// Set all the query parameters for the request
func setParams(req *http.Request, params map[string]string) {
	if params != nil {
		query := req.URL.Query()

		for key, val := range params {
			query.Add(key, val)
		}

		req.URL.RawQuery = query.Encode()
	}
}

// Set Basic auth for request
func setBasicAuth(req *http.Request, username, password string) {
	req.SetBasicAuth(username, password)
}

// Set Headers here for requests
func setRequestHeaders(req *http.Request, headers map[string]string) {
	for key, val := range headers {
		req.Header.Set(key, val)
	}

	req.Header.Set("Content-Type", "application/json")
}

// Get http client with timeout set
func getHttpClient() *http.Client {

	//TODO: Monitor this timeout.  If we start to get backed up with slow responses turn this down
	// Otherwise this may be fine.
	timeout := time.Second * 30

	client := &http.Client{Timeout: timeout}

	return client
}
