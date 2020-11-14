package stats

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
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
	hostname   string
}

// GetData makes the reqest to get the data and returns them.
func (sc SimpleClient) GetData() (Data, error) {
	res, err := sc.makeRequest()
	if err != nil {
		return Data{}, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Data{}, err
	}

	log.Println(string(body))
	var jd []jsonData
	err = json.Unmarshal(body, &jd)
	if len(jd) < 1 {
		return Data{}, errors.New("No data received from the db")
	}
	return jd[0].transformToData()
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

	baseURL.Path += fmt.Sprintf("/hosts/%s/stats", sc.hostname)
	q := baseURL.Query()
	q.Add("limit", "1")
	baseURL.RawQuery = q.Encode()

	return baseURL.String(), nil
}

func (sc SimpleClient) buildRequest() (*http.Request, error) {
	url, err := sc.buildURL()
	if err != nil {
		return &http.Request{}, err
	}
	log.Println(url)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return &http.Request{}, err
	}
	request.Header.Add("Token", sc.token)
	return request, nil
}

func (sc SimpleClient) makeRequest() (*http.Response, error) {
	request, err := sc.buildRequest()
	if err != nil {
		return &http.Response{}, err
	}
	return sc.httpClient.Do(request)
}

// NewSimpleClient is a convenient constructor for the SimpleClient with the http.DefaultClient.
func NewSimpleClient(token string, url string, hostname string) SimpleClient {
	return SimpleClient{http.DefaultClient, token, url, hostname}
}
