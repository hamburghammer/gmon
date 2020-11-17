package main

import (
	"io"
	"os"

	"github.com/hamburghammer/gmon/alert"
	"github.com/hamburghammer/gmon/config"
	"github.com/hamburghammer/gmon/stats"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

var logPackage = log.WithField("package", "main")

type arguments struct {
	ConfigPath string `short:"c" long:"config" default:"./config.toml" description:"Set the path to the configuration file." env:"GMON_CONFIG_PATH"`
	RulePath   string `short:"r" long:"rules" default:"./rules.toml" description:"Set the path to the file with the rules." env:"GMON_RULE_PATH"`
	Verbose    bool   `short:"v" long:"verbose" description:"Set the logging output level to trace."`
	Quiet      bool   `short:"q" long:"quiet" description:"Set the logging output level to error."`
	JSON       bool   `long:"json" description:"Set the logging format to json"`
}

func parseArgs() arguments {
	var args = arguments{}
	_, err := flags.Parse(&args)
	if err != nil {
		os.Exit(1)
	}
	return args
}

func initLogging(args arguments) {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	if args.Verbose {
		log.SetLevel(log.TraceLevel)
	}
	if args.Quiet {
		log.SetLevel(log.ErrorLevel)
	}
	if args.JSON {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func main() {
	args := parseArgs()

	initLogging(args)

	configReader, err := loadFile(args.ConfigPath)
	if err != nil {
		logPackage.Fatalln(err)
	}
	configuration, err := config.NewTomlConfigLoader(configReader).Load()
	if err != nil {
		logPackage.Fatalf("Parsing the toml file '%s' produced an error: %s\n", args.ConfigPath, err)
	}
	logPackage.Println(configuration)

	rulesReader, err := loadFile(args.RulePath)
	if err != nil {
		logPackage.Fatalln(err)
	}
	rules, err := config.NewTOMLRulesLoader(rulesReader).Load()
	if err != nil {
		logPackage.Fatalf("Parsing the toml file '%s' produced an error: %s\n", args.RulePath, err)
	}

	statsClient := stats.NewSimpleClient(configuration.Stats.Token, configuration.Stats.Endpoint, configuration.Stats.Hostname)
	gotifyClient := alert.NewGotifyClient(configuration.Gotify.Token, configuration.Gotify.Endpoint)

	monitoring := NewMonitoring(statsClient, gotifyClient, rules, configuration.Interval)
	err = monitoring.Monitor()
	if err != nil {
		logPackage.Fatalln(err)
	}
}

func loadFile(path string) (io.Reader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}
