package config

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockFile struct {
	content   []byte
	readIndex int
}

func (f *mockFile) Read(b []byte) (n int, err error) {
	if f.readIndex >= len(f.content) {
		err = io.EOF
		return
	}
	n = copy(b, f.content[f.readIndex:])
	f.readIndex += n
	return
}

func TestTomlFileLoaderLoad(t *testing.T) {
	file := &mockFile{content: []byte(`[stats]
	endpoint = "https://stats.example.com"
	hostname = "test"
	token = "xxx"
	
	[gotify]
	endpoint = "https://push.example.com"
	token = "xxx"`)}

	statsConfig := Stats{Endpoint: "https://stats.example.com", Hostname: "test", Token: "xxx"}
	gotifyConfig := Gotify{Endpoint: "https://push.example.com", Token: "xxx"}
	want := Data{Stats: statsConfig, Gotify: gotifyConfig}
	got, err := TomlLoader{file: file}.Load()

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
