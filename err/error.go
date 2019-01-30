package err

type errorString struct {
        msg string
}

func (e *errorString) Error() string {
        return e.msg
}

func NewError(msg string) error {
        return &errorString{msg}
}
