package slack

type Error struct {
	Message string
}

func (err *Error) Error() string {
	return err.Message
}
