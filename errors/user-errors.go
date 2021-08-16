package errors

type UserError string

const (
	ErrUserNotFound UserError = "User not found"
)

func (u *UserError) Message() string {
	return string(*u)
}
