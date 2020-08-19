package stats

import (
	"encoding/json"
	"io/ioutil"
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

	var jd jsonData
	err = json.Unmarshal(body, &jd)
	return jd.transformToData()
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

	queryString := baseURL.Query()
	queryString.Add("token", sc.token)

	baseURL.RawQuery = queryString.Encode()
	baseURL.Path += sc.hostname

	return baseURL.String(), nil
}

func (sc SimpleClient) buildRequest() (*http.Request, error) {
	url, err := sc.buildURL()
	if err != nil {
		return &http.Request{}, err
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return &http.Request{}, err
	}
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
