package config

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"
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

func TestTomlLoader_LoadStatsConfig(t *testing.T) {
	file := &mockReader{content: []byte(`[stats]
	endpoint = "https://stats.example.com"
	hostname = "test"
	token = "xxx"`)}

	statsConfig := Stats{Endpoint: "https://stats.example.com", Hostname: "test", Token: "xxx"}
	want := Data{Stats: statsConfig}
	got, err := TomlLoader{reader: file}.Load()

	require.Nil(t, err)
	require.Equal(t, want, got)
}

func TestTomlLoader_LoadGotifyConfig(t *testing.T) {
	file := &mockReader{content: []byte(`[gotify]
	endpoint = "https://push.example.com"
	token = "xxx"`)}

	gotifyConfig := Gotify{Endpoint: "https://push.example.com", Token: "xxx"}
	want := Data{Gotify: gotifyConfig}
	got, err := TomlLoader{reader: file}.Load()

	require.Nil(t, err)
	require.Equal(t, want, got)
}
