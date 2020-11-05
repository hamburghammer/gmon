package alert

import (
	"bytes"
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

func getMockClientWithValidResponse() mockHTTPClient {
	responseJSON := ""
	return mockHTTPClient{response: responseJSON, statusCode: 200}
}

func TestSimpleClientRequestBuilding(t *testing.T) {

	data := Data{
		Title:    "title",
		Message:  "message",
		Priority: "high",
	}

	t.Run("should build url with the token as header", func(t *testing.T) {
		mockClient := getMockClientWithValidResponse()
		token := "xxx"
		simpleClient := NewSimpleClient(token, "https://example.com")
		simpleClient = simpleClient.WithHTTPClient(&mockClient)

		_, err := simpleClient.Notify(data)
		assert.Nil(t, err)

		assert.Equal(t, token, mockClient.req.Header.Get("X-Gotify-Key"))
	})

	t.Run("should return error if url could not be parse", func(t *testing.T) {
		mockClient := &mockHTTPClient{}
		token := "xxx"
		simpleClient := NewSimpleClient(token, " https:")
		simpleClient = simpleClient.WithHTTPClient(mockClient)

		_, err := simpleClient.Notify(data)
		assert.NotNil(t, err)

		want := "parse \" https:\": first path segment in URL cannot contain colon"

		assert.Equal(t, want, err.Error())
	})
}
