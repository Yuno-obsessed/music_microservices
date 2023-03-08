package mail

type Subject string

func (s Subject) Text() string {
	return string(s)
}

const (
	QrCode        Subject = "Your Qr code ticket for a musical event"
	Registration  Subject = "Registration successful"
	Login         Subject = "You were logged in"
	PaymentFailed Subject = "Payment for a musical event failed"
	// New event from your subscriptions
)
