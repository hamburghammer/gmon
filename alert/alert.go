package alert

import log "github.com/sirupsen/logrus"

var logPackage = log.WithField("package", "alert")

// Notifier to send notifications to an endpoint
type Notifier interface {
	Notify(Data) error
}

// Data hold the information to send a new notification
type Data struct {
	Title    string `json:"title"`
	Message  string `json:"message"`
	Priority int    `json:"priority"`
}
