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
		log.Fatalf("Parsing the toml file '%s' produced an error: %s", args.RulePath, err)
	}
	log.Printf("%+v", rules)

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
		data, err := statsClient.GetData()
		if err != nil {
			return err
		}

		var f analyse.Analyser
		f = analyse.CPURule{}
		f.Analyse(stats.Data{})

		applyRules(cpuRulesToAnalyse(rules.CPU), data, gotifyClient)
		applyRules(diskRulesToAnalyse(rules.Disk), data, gotifyClient)
		applyRules(ramRulesToAnalyse(rules.RAM), data, gotifyClient)

		time.Sleep(time.Duration(interval) * time.Minute)
	}
}

func applyRules(rules []analyse.Analyser, stat stats.Data, notifier alert.Notifier) error {
	for _, rule := range rules {
		err := applyRule(rule, stat, notifier)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	return nil
}

func applyRule(rule analyse.Analyser, stat stats.Data, notifier alert.Notifier) error {
	if rule.IsDeactivated() {
		return nil
	}

	result, err := rule.Analyse(stat)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(result)

	if result.AlertStatus != analyse.StatusOK {
		err = notifier.Notify(alert.Data{Title: result.Title, Message: result.StatusMessage})
		if err != nil {
			return err
		}
	}

	return nil
}

func cpuRulesToAnalyse(rs []analyse.CPURule) []analyse.Analyser {
	rules := make([]analyse.Analyser, len(rs))
	for i, r := range rs {
		rules[i] = r
	}

	return rules
}

func diskRulesToAnalyse(rs []analyse.DiskRule) []analyse.Analyser {
	rules := make([]analyse.Analyser, len(rs))
	for i, r := range rs {
		rules[i] = r
	}

	return rules
}

func ramRulesToAnalyse(rs []analyse.RAMRule) []analyse.Analyser {
	rules := make([]analyse.Analyser, len(rs))
	for i, r := range rs {
		rules[i] = r
	}

	return rules
}
