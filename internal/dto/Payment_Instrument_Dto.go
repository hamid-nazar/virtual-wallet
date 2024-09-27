package dto

const (
	InstrumentTypeCard InstrumentType = iota
	InstrumentTypeWallet
)

func (t InstrumentType) String() string {
	// return [...]string{"Card", "Wallet"}[t]
	switch t {
	case InstrumentTypeCard:
		return "CARD"
	case InstrumentTypeWallet:
		return "WALLET"
	default:
		return "Unknown"
	}
}

type NewPaymentInstrumentDto struct {
	Name           string         `json:"name"`
	OwnerId        string         `json:"ownerId"`
	InstrumentType InstrumentType `json:"instrumentType"`
}
