package exception

type EmailError struct {
	Message string
}

func (emailError EmailError) Error() string {
	return emailError.Message
}
