package alert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// HTTPClient interface exist to exchange the request.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// SimpleClient is a client implementation to get the data from an http endpoint.
type SimpleClient struct {
	httpClient HTTPClient
	token      string
	url        string
}

// Notify makes the request to push the data to the gotify server.
func (sc SimpleClient) Notify(data Data) error {
	_, err := sc.makeRequest(data)
	if err != nil {
		return err
	}
	return nil
}

// WithHTTPClient sets a new HttpClient and returns a new SimpleClient.
func (sc SimpleClient) WithHTTPClient(httpClient HTTPClient) SimpleClient {
	sc.httpClient = httpClient
	return sc
}

func (sc SimpleClient) buildURL() (string, error) {
	baseURL, err := url.Parse(sc.url)
	if err != nil {
		return "", err
	}

	baseURL.Path = "/message"
	return baseURL.String(), nil
}

func (sc SimpleClient) buildRequest(data Data) (*http.Request, error) {
	url, err := sc.buildURL()
	if err != nil {
		return &http.Request{}, err
	}

	jsonStr, err := json.Marshal(data)
	if err != nil {
		return &http.Request{}, fmt.Errorf("Mashaling the data produced an error: %w", err)
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return &http.Request{}, err
	}

	request.Header.Add("X-Gotify-Key", sc.token)
	request.Header.Add("Content-Type", "application/json")
	return request, nil
}

func (sc SimpleClient) makeRequest(data Data) (*http.Response, error) {
	request, err := sc.buildRequest(data)
	if err != nil {
		return &http.Response{}, err
	}
	return sc.httpClient.Do(request)
}

// NewSimpleClient is a convenient constructor for the SimpleClient with the http.DefaultClient.
func NewSimpleClient(token string, url string) SimpleClient {
	return SimpleClient{http.DefaultClient, token, url}
}
