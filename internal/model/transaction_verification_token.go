package model

import "time"

type TransactionVerificationToken struct {
	id          string
	token       string
	transaction Transaction
	createdAt   time.Time
	expiryDate  time.Time
}

func NewTransactionVerificationToken(id, token string, transaction Transaction) *TransactionVerificationToken {
	return &TransactionVerificationToken{
		id:          id,
		token:       token,
		transaction: transaction,
		createdAt:   time.Now(),
		expiryDate:  time.Now().Add(5 * time.Minute),
	}
}

func (t *TransactionVerificationToken) getExpiryDate() time.Time {
	return t.expiryDate
}

func (t *TransactionVerificationToken) GetId() string {
	return t.id
}

func (t *TransactionVerificationToken) GetToken() string {
	return t.token
}

func (t *TransactionVerificationToken) GetTransaction() Transaction {
	return t.transaction
}

func (t *TransactionVerificationToken) GetCreatedAt() time.Time {
	return t.createdAt
}

func (t *TransactionVerificationToken) GetExpiryDate() time.Time {
	return t.expiryDate
}

func (t *TransactionVerificationToken) SetExpiryDate(expiryDate time.Time) {
	t.expiryDate = expiryDate
}

func (t *TransactionVerificationToken) SetId(id string) {
	t.id = id
}

func (t *TransactionVerificationToken) SetToken(token string) {
	t.token = token
}

func (t *TransactionVerificationToken) SetTransaction(transaction Transaction) {
	t.transaction = transaction
}

func (t *TransactionVerificationToken) SetCreatedAt(createdAt time.Time) {
	t.createdAt = createdAt
}
