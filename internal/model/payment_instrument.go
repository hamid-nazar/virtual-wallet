package model

type InstrumentType int

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

type PaymentInstrument struct {
	id             string
	owner          User
	instrumentType InstrumentType
	name           string
}

func NewPaymentInstrument(id string, owner User, instrumentType InstrumentType, name string) *PaymentInstrument {
	return &PaymentInstrument{
		id:             id,
		owner:          owner,
		instrumentType: instrumentType,
		name:           name,
	}
}

func (p *PaymentInstrument) GetId() string {
	return p.id
}

func (p *PaymentInstrument) GetOwner() User {
	return p.owner
}

func (p *PaymentInstrument) GetInstrumentType() InstrumentType {
	return p.instrumentType
}

func (p *PaymentInstrument) GetName() string {
	return p.name
}

func (p *PaymentInstrument) SetId(id string) {
	p.id = id
}

func (p *PaymentInstrument) SetOwner(owner User) {
	p.owner = owner
}

func (p *PaymentInstrument) SetInstrumentType(instrumentType InstrumentType) {
	p.instrumentType = instrumentType
}

func (p *PaymentInstrument) SetName(name string) {
	p.name = name
}
