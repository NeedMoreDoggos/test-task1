package creatres

type Response struct {
	ID      string `json:"status"`
	Balance string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func OK(id, balance string) Response {
	return Response{
		ID:      id,
		Balance: balance,
	}
}

func Error(msg string) Response {
	return Response{
		Error: msg,
	}
}
