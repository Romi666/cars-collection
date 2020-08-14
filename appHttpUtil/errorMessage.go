package appHttpUtil

type ErrorResponseBuilder interface {
	Code() int
	Message() string
}

type ErrorMessage struct{
	ErrorID int `json:"errorId"`
	Msg string `json:"message"`
}

func (e ErrorMessage) Code() int {
	return e.ErrorID
}

func (e ErrorMessage) Message() string {
	return e.Msg
}

func NewErrorMessage(code int, message string) ErrorResponseBuilder{
	return &ErrorMessage{
		ErrorID: code,
		Msg: message,
	}
}