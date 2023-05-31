package response

type Error struct {
	Msg        string `json:"msg"`
	StatusCode int    `json:"statusCode"`
}

func NewError(msg string, statusCode int) *Error {
	return &Error{
		Msg:        msg,
		StatusCode: statusCode,
	}
}
