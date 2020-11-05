package alert

// Notifier to send notifications to an endpoint
type Notifier interface {
	Notify(Data) error
}

// Data hold the information to send a new notification
type Data struct {
	Title    string
	Message  string
	Priority string
}
