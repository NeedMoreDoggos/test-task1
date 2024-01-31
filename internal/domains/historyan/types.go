package historyan

import "github.com/NeedMoreDoggos/test-task1/internal/domains"

type Response struct {
	Error  string          `json:"error,omitempty"`
	Events []domains.Event `json:"events"`
	OK     bool            `json:"ok"`
}

func Error(msg string) Response {
	return Response{
		OK:    false,
		Error: msg,
	}
}

func OK(events []domains.Event) Response {
	return Response{
		OK:     true,
		Events: events,
	}
}
