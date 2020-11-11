package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/hamburghammer/gmon/alert"
	"github.com/hamburghammer/gmon/analyse"
	"github.com/hamburghammer/gmon/config"
	"github.com/hamburghammer/gmon/stats"
	"github.com/jessevdk/go-flags"
)

type arguments struct {
	ConfigPath string `short:"c" long:"config" default:"./config.toml" description:"Set the path to the configuration file." env:"GMON_CONFIG_PATH"`
	RulePath   string `short:"r" long:"rules" default:"./rules.toml" description:"Set the path to the file with the rules." env:"GMON_RULE_PATH"`
}

func parseArgs() arguments {
	var args = arguments{}
	_, err := flags.Parse(&args)
	if err != nil {
		os.Exit(1)
	}
	return args
}

func main() {
	args := parseArgs()

	configReader, err := loadFile(args.ConfigPath)
	if err != nil {
		log.Fatalln(err)
	}
	configuration, err := config.NewTomlConfigLoader(configReader).Load()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(configuration)

	rulesReader, err := loadFile(args.RulePath)
	if err != nil {
		log.Fatalln(err)
	}
	rules, err := config.NewTOMLRulesLoader(rulesReader).Load()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(rules)

	statsClient := stats.NewSimpleClient(configuration.Stats.Token, configuration.Stats.Endpoint, configuration.Stats.Hostname)
	gotifyClient := alert.NewGotifyClient(configuration.Gotify.Token, configuration.Gotify.Endpoint)

	err = Monitor(statsClient, gotifyClient, configuration.Interval, rules)
	if err != nil {
		log.Fatalln(err)
	}
}

func loadFile(path string) (io.Reader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// Monitor the stats -> get, analyse and notify
func Monitor(statsClient stats.Client, gotifyClient alert.Notifier, interval int, rules config.Rules) error {
	for {
		stats, err := statsClient.GetData()
		if err != nil {
			return err
		}

		for _, rule := range rules.CPU {
			result, err := rule.Analyse(stats)
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println(result)
			if result.AlertStatus != analyse.StatusOK {
				err = gotifyClient.Notify(alert.Data{Title: result.Title, Message: result.StatusMessage})
				if err != nil {
					return err
				}
			}
		}
		time.Sleep(time.Duration(interval) * time.Minute)
	}
}
