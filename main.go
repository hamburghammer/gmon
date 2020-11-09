package main

import (
	"encoding/json"
	"log"

	"github.com/hamburghammer/gmon/alert"
	"github.com/hamburghammer/gmon/stats"
)

func main() {
	statsClient := stats.NewSimpleClient("foo", "http://localhost:8080", "foo")
	data, err := statsClient.GetData()
	if err != nil {
		log.Println(err)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	gotifyClient := alert.NewGotifyClient("AzCkehMSkHFlphf", "http://localhost:80")
	gotifyClient.Notify(alert.Data{Title: "Host: foo", Message: string(jsonData)})
}
