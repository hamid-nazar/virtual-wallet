package model

import (
	"time"
)

type transactionType int

const (
	SmallAmount transactionType = iota
	LargeVerified
	LargeUnverified
	CardToWallet
	Donation
)

func (t transactionType) String() string {
	// return [...]string{"SmallAmount", "LargeVerified", "LargeUnverified", "CardToWallet", "Donation"}[t]
	switch t {
	case SmallAmount:
		return "SMALL_AMOUNT"
	case LargeVerified:
		return "LARGE_VERIFIED"
	case LargeUnverified:
		return "LARGE_UNVERIFIED"
	case CardToWallet:
		return "CARD_TO_WALLET"
	case Donation:
		return "DONATION"
	default:
		return "Unknown"
	}
}

type Transaction struct {
	id                 string
	timeStamp          time.Time
	amount             float64
	senderInstrument   PaymentInstrument
	receiverInstrument PaymentInstrument
	description        string
	transactionType    transactionType
	withDonation       bool
}

func NewTransaction(amount float64, senderInstrument PaymentInstrument, receiverInstrument PaymentInstrument, description string, transactionType transactionType, withDonation bool) *Transaction {
	return &Transaction{
		amount:             amount,
		senderInstrument:   senderInstrument,
		receiverInstrument: receiverInstrument,
		description:        description,
		transactionType:    transactionType,
		withDonation:       withDonation,
	}
}

func (t *Transaction) GetId() string {
	return t.id
}

func (t *Transaction) GetTimeStamp() time.Time {
	return t.timeStamp
}

func (t *Transaction) GetAmount() float64 {
	return t.amount
}

func (t *Transaction) GetSenderInstrument() PaymentInstrument {
	return t.senderInstrument
}

func (t *Transaction) GetReceiverInstrument() PaymentInstrument {
	return t.receiverInstrument
}

func (t *Transaction) GetDescription() string {
	return t.description
}

func (t *Transaction) GetTransactionType() transactionType {
	return t.transactionType
}

func (t *Transaction) GetWithDonation() bool {
	return t.withDonation
}

func (t *Transaction) SetId(id string) {
	t.id = id
}

func (t *Transaction) SetTimeStamp(timeStamp time.Time) {
	t.timeStamp = timeStamp
}

func (t *Transaction) SetAmount(amount float64) {
	t.amount = amount
}

func (t *Transaction) SetSenderInstrument(senderInstrument PaymentInstrument) {
	t.senderInstrument = senderInstrument
}

func (t *Transaction) SetReceiverInstrument(receiverInstrument PaymentInstrument) {
	t.receiverInstrument = receiverInstrument
}

func (t *Transaction) SetDescription(description string) {
	t.description = description
}

func (t *Transaction) SetTransactionType(transactionType transactionType) {
	t.transactionType = transactionType
}

func (t *Transaction) SetWithDonation(withDonation bool) {
	t.withDonation = withDonation
}
