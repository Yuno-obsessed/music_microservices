package consts

type PaymentMethod string

const (
	PAYPAL PaymentMethod = "pay-pal"
	CRYPTO PaymentMethod = "crypto"
	GOOGLE PaymentMethod = "google"
)

func GetPaymentMethods() []PaymentMethod {
	return []PaymentMethod{
		PAYPAL,
		CRYPTO,
		GOOGLE,
	}
}
