package mail

type MessageType string

func (m MessageType) Text() string {
	return string(m)
}

var (
	SuccessfulRegistration    MessageType = "Your Account was successfully registered"
	SuccessfulLogin           MessageType = "Successful sign-in using your email"
	NewEventFromSubscriptions MessageType = "Come and check a new event from one of your subscriptions!"
)
