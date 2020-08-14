package stats

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

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

func TestSimpleClientRequestBuilding(t *testing.T) {
	t.Run("should build url with the hostname as path", func(t *testing.T) {
		mockClient := &mockHTTPClient{}
		hostname := "yyy"
		simpleClient := NewSimpleClient("xxx", "https://example.com", hostname)
		simpleClient = simpleClient.WithHTTPClient(mockClient)

		_, err := simpleClient.GetData()
		assert.Nil(t, err)

		want := fmt.Sprintf("/%s", hostname)

		assert.Equal(t, want, mockClient.req.URL.Path)
	})

	t.Run("should build url with the token as query encoded", func(t *testing.T) {
		mockClient := &mockHTTPClient{}
		token := "xxx"
		simpleClient := NewSimpleClient(token, "https://example.com", "yyy")
		simpleClient = simpleClient.WithHTTPClient(mockClient)

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
