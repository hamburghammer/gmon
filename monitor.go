package main

import (
	"context"
	"time"

	"github.com/hamburghammer/gmon/alert"
	"github.com/hamburghammer/gmon/analyse"
	"github.com/hamburghammer/gmon/config"
	"github.com/hamburghammer/gmon/stats"
)

// NewMonitoring a constructor for the Monitoring struct.
func NewMonitoring(statsClient stats.Client, notifierClient alert.Notifier, rules config.Rules, interval int) Monitoring {
	return Monitoring{statsClient: statsClient, notifierClient: notifierClient, rules: rules, interval: interval, ctx: context.Background()}
}

// Monitoring hold all the configuration to monitor a service.
type Monitoring struct {
	statsClient    stats.Client
	notifierClient alert.Notifier
	rules          config.Rules
	interval       int
	ctx            context.Context
}

// WithContext changes the context of the monitoring struct and returns it.
func (m Monitoring) WithContext(ctx context.Context) Monitoring {
	m.ctx = ctx
	return m
}

// Monitor the stats -> get, analyse and notify.
// This method can be canceled over the context and can be run in a goroutine if the used
// clients support shared usage.
// It currently monitors CPU, RAM and Disk usage with the given rule set.
// Returns an error if a client operation fails
func (m Monitoring) Monitor() error {
	for {
		select {
		case <-m.ctx.Done():
			logPackage.Warnln(m.ctx.Err())
			return nil
		default:
			data, err := m.statsClient.GetData()
			if err != nil {
				return err
			}

			var f analyse.Analyser
			f = analyse.CPURule{}
			f.Analyse(stats.Data{})

			m.applyRules(m.rules.GetCPU(), data)
			if err != nil {
				return err
			}
			m.applyRules(m.rules.GetDisk(), data)
			if err != nil {
				return err
			}
			m.applyRules(m.rules.GetRAM(), data)
			if err != nil {
				return err
			}

			time.Sleep(time.Duration(m.interval) * time.Minute)
		}
	}
}

func (m Monitoring) applyRules(rules []analyse.Analyser, stat stats.Data) error {
	for _, rule := range rules {
		err := m.applyRule(rule, stat)
		if err != nil {
			logPackage.Errorln(err)
			continue
		}
	}

	return nil
}

func (m Monitoring) applyRule(rule analyse.Analyser, stat stats.Data) error {
	if rule.IsDeactivated() {
		return nil
	}

	result, err := rule.Analyse(stat)
	if err != nil {
		return err
	}
	logPackage.Infoln(result)

	if result.AlertStatus != analyse.StatusOK {
		err = m.notifierClient.Notify(alert.Data{Title: result.Title, Message: result.StatusMessage})
		if err != nil {
			return err
		}
	}

	return nil
}
