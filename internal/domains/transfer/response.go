package transfer

type Response struct {
	Error string `json:"error,omitempty"`
	OK    bool   `json:"ok"`
}

func Error(msg string) Response {
	return Response{
		OK:    false,
		Error: msg,
	}
}

func OK() Response {
	return Response{
		OK: true,
	}
}
