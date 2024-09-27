package model

type Wallet struct {
	id      string
	name    string
	balance float64
	deleted bool
}

func NewWallet(id, name string, balance float64) *Wallet {
	return &Wallet{
		id:      id,
		name:    name,
		balance: balance,
		deleted: false,
	}
}

func (w *Wallet) GetId() string {
	return w.id
}

func (w *Wallet) GetName() string {
	return w.name
}

func (w *Wallet) GetBalance() float64 {
	return w.balance
}

func (w *Wallet) GetDeleted() bool {
	return w.deleted
}

func (w *Wallet) SetName(name string) {
	w.name = name
}

func (w *Wallet) SetBalance(balance float64) {
	w.balance = balance
}

func (w *Wallet) SetDelete(deleted bool) {
	w.deleted = deleted
}
