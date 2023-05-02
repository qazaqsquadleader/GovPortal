package errormsg

type UserError struct {
	msg    string
	status int
}

func NewUserError(msg string, status int) error {
	return &UserError{msg, status}
}

func (e *UserError) Error() string {
	return e.msg
}

func (e *UserError) Status() int {
	return e.status
}
