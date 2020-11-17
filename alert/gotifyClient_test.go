package alert

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
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

func TestGotifyClientRequestBuilding(t *testing.T) {

	data := Data{
		Title:    "title",
		Message:  "message",
		Priority: 0,
	}

	t.Run("should build url with the token as header", func(t *testing.T) {
		mockClient := getMockClientWithValidResponse()
		token := "xxx"
		simpleClient := NewGotifyClient(token, "https://example.com")
		simpleClient = simpleClient.WithHTTPClient(&mockClient)

		err := simpleClient.Notify(data)
		require.Nil(t, err)

		require.Equal(t, token, mockClient.req.Header.Get("X-Gotify-Key"))
	})

	t.Run("should return error if url could not be parse", func(t *testing.T) {
		mockClient := &mockHTTPClient{}
		token := "xxx"
		simpleClient := NewGotifyClient(token, " https:")
		simpleClient = simpleClient.WithHTTPClient(mockClient)

		err := simpleClient.Notify(data)
		require.NotNil(t, err)

		want := "parse \" https:\": first path segment in URL cannot contain colon"

		require.Equal(t, want, err.Error())
	})
}
