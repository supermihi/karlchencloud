package server

type CloudError struct {
	msg string
}

func Error(msg string) CloudError {
	return CloudError{msg}
}

func (c CloudError) Error() string {
	return c.msg
}
