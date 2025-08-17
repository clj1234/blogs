package myErrors

type ResponseError struct {
	ErrCode int64
	Message string
}

func (res *ResponseError) Error() string {
	if res == nil {
		return ""
	}
	return res.Message
}

func NewResponseError(text string) *ResponseError {
	return &ResponseError{Message: text}
}
