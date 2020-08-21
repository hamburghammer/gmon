package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTomlLoaderLoad(t *testing.T) {
	data := []byte(`
	[stats]
	endpoint = "https://stats.example.com"
	hostname = "test"
	token = "xxx"
	
	[gotify]
	endpoint = "https://push.example.com"
	token = "xxx"
	`)

	statsConfig := Stats{Endpoint: "https://stats.example.com", Hostname: "test", Token: "xxx"}
	gotifyConfig := Gotify{Endpoint: "https://push.example.com", Token: "xxx"}
	want := Data{Stats: statsConfig, Gotify: gotifyConfig}
	got, err := tomlLoader{toml: &data}.Load()

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestTomlFileLoaderLoad(t *testing.T) {
	file, err := os.Open("./test_files/config.toml")
	if err != nil {
		t.Fatal("test file not found")
	}

	statsConfig := Stats{Endpoint: "https://stats.example.com", Hostname: "test", Token: "xxx"}
	gotifyConfig := Gotify{Endpoint: "https://push.example.com", Token: "xxx"}
	want := Data{Stats: statsConfig, Gotify: gotifyConfig}
	got, err := TomlFileLoader{file: file}.Load()

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
