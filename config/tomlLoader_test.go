package config

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockReader struct {
	content   []byte
	readIndex int
}

func (mr *mockReader) Read(b []byte) (n int, err error) {
	if mr.readIndex >= len(mr.content) {
		err = io.EOF
		return
	}
	n = copy(b, mr.content[mr.readIndex:])
	mr.readIndex += n
	return
}

func TestTomlFileLoaderLoad(t *testing.T) {
	file := &mockReader{content: []byte(`[stats]
	endpoint = "https://stats.example.com"
	hostname = "test"
	token = "xxx"
	
	[gotify]
	endpoint = "https://push.example.com"
	token = "xxx"`)}

	statsConfig := Stats{Endpoint: "https://stats.example.com", Hostname: "test", Token: "xxx"}
	gotifyConfig := Gotify{Endpoint: "https://push.example.com", Token: "xxx"}
	want := Data{Stats: statsConfig, Gotify: gotifyConfig}
	got, err := TomlLoader{reader: file}.Load()

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
