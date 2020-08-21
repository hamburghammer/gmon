package stats

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockHTTPClient struct {
	response   string
	statusCode int
	req        http.Request
}

func (mhc *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	mhc.req = *req
	response := http.Response{
		Body:       ioutil.NopCloser(bytes.NewBufferString(mhc.response)),
		StatusCode: mhc.statusCode,
	}
	return &response, nil
}

func newSimpleClientWithMockResponseHTTPClient(response string, statusCode int) Client {
	simpleClient := NewSimpleClient("xxx", "https://example.com", "yyy")
	return simpleClient.WithHTTPClient(&mockHTTPClient{response: response, statusCode: statusCode})
}

func getMockClientWithValidResponse() mockHTTPClient {
	responseJSON := "{\"Date\":\"2020-08-19T17:45:56+02:00\",\"CPU\":1.9900497512574382,\"Mem\":\"4621/16022\",\"Disk\":\"51271/224323\",\"Processes\":[{\"Name\":\"gstat\",\"Pid\":1,\"CPU\":37.58064430461327}]}"
	return mockHTTPClient{response: responseJSON, statusCode: 200}
}

func TestSimpleClientRequestBuilding(t *testing.T) {
	t.Run("should build url with the hostname as path", func(t *testing.T) {
		mockClient := getMockClientWithValidResponse()
		hostname := "yyy"
		simpleClient := NewSimpleClient("xxx", "https://example.com", hostname)
		simpleClient = simpleClient.WithHTTPClient(&mockClient)

		_, err := simpleClient.GetData()
		assert.Nil(t, err)

		want := fmt.Sprintf("/%s", hostname)

		assert.Equal(t, want, mockClient.req.URL.Path)
	})

	t.Run("should build url with the token as query encoded", func(t *testing.T) {
		mockClient := getMockClientWithValidResponse()
		token := "xxx"
		simpleClient := NewSimpleClient(token, "https://example.com", "yyy")
		simpleClient = simpleClient.WithHTTPClient(&mockClient)

		_, err := simpleClient.GetData()
		assert.Nil(t, err)

		want := fmt.Sprintf("token=%s", token)

		assert.Equal(t, want, mockClient.req.URL.RawQuery)
	})

	t.Run("should return error if url could not be parse", func(t *testing.T) {
		mockClient := &mockHTTPClient{}
		token := "xxx"
		simpleClient := NewSimpleClient(token, " https:", "yyy")
		simpleClient = simpleClient.WithHTTPClient(mockClient)

		_, err := simpleClient.GetData()
		assert.NotNil(t, err)

		want := "parse \" https:\": first path segment in URL cannot contain colon"

		assert.Equal(t, want, err.Error())
	})
}

func TestSimpleClientGetData(t *testing.T) {
	t.Run("should deserialize the data", func(t *testing.T) {
		responseJSON := "{\"Date\":\"2020-08-19T17:45:56+02:00\",\"CPU\":1.9900497512574382,\"Mem\":\"4621/16022\",\"Disk\":\"51271/224323\",\"Processes\":[{\"Name\":\"gstat\",\"Pid\":1,\"CPU\":37.58064430461327}]}"
		parsedTime, _ := time.Parse(time.RFC3339, "2020-08-19T17:45:56+02:00")

		mockClient := newSimpleClientWithMockResponseHTTPClient(responseJSON, 200)

		want := Data{Date: parsedTime, CPU: 1.9900497512574382, Mem: Memory{Used: 4621, Total: 16022}, Disk: Memory{Used: 51271, Total: 224323}, Processes: []Process{{Name: "gstat", Pid: 1, CPU: 37.58064430461327}}}
		got, err := mockClient.GetData()

		assert.Nil(t, err)
		assert.Equal(t, want, got)
	})
}
