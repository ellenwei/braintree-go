package braintree

type CreditCard struct {
	Token          string             `xml:"token,omitempty"`
	Number         string             `xml:"number"`
	ExpirationDate string             `xml:"expiration-date"`
	CVV            string             `xml:"cvv,omitempty"`
	Options        *CreditCardOptions `xml:"options,omitempty"`
}

type CreditCardOptions struct {
	VerifyCard bool `xml:"verify-card,omitempty"`
}