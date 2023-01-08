package entity

type EventId string

type Event struct {
	Id           EventId
	Name         string
	Detail       string
	Location     string
	State        string
	Type         string
}

type State string

var (
	OPEN_STATE   State = "open"
	CLOSE_STATE  State = "close"
	CANCEL_STATE State = "cancel"
)

type Type string

var (
	ONCE_TYPE Type = "once"
)
