package main

import (
	"encoding/json"
	"log"

	"github.com/hamburghammer/gmon/alert"
	"github.com/hamburghammer/gmon/analyse"
	"github.com/hamburghammer/gmon/stats"
)

func main() {
	statsClient := stats.NewSimpleClient("foo", "http://localhost:8080", "foo")
	data, err := statsClient.GetData()
	if err != nil {
		log.Println(err)
	}

	rules := []analyse.Analyser{&analyse.CPURule{Rule: analyse.Rule{Compare: ">", Name: "Foo cpu"}, Warning: 0.5, Alert: 1}}

	for _, rule := range rules {
		alert, err := rule.Analyse(data)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(alert)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	gotifyClient := alert.NewGotifyClient("AzCkehMSkHFlphf", "http://localhost:80")
	gotifyClient.Notify(alert.Data{Title: "Host: foo", Message: string(jsonData)})
}
