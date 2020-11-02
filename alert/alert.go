package alert

// the interface to be implemented by the evaluation module to pass data that needs to be send via api
type NotifyInterface interface {
	Notify(Data) (error)
}

type Data struct {
	Title string
	Message string
	Priority string
}


