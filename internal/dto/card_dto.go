package dto

type NewCardDto struct {
	NewPaymentInstrumentDto NewPaymentInstrumentDto `json:"newPaymentInstrumentDto"`
	Number                  string                  `json:"number,omitempty"`
	CardHolder              string                  `json:"cardHolder,omitempty"`
	ExpirationDate          string                  `json:"expirationDate,omitempty"`
	CVV                     string                  `json:"cvv,omitempty"`
}
type CardDetailsDTO struct {
	Number string `json:"number,omitempty"`
	Name   string `json:"name,omitempty"`
	CVV    string `json:"cvv,omitempty"`
}
