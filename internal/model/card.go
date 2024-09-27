package model

import (
	"time"
)

type Card struct {
	id             string
	name           string
	holderName     string
	number         string
	expirationDate time.Time
	cvv            string
	deleted        bool
}

func NewCard(name, holderName, number string, expirationDate time.Time, cvv string) *Card {
	return &Card{
		id:             uuid(),
		name:           name,
		holderName:     holderName,
		number:         number,
		expirationDate: expirationDate,
		cvv:            cvv,
		deleted:        false,
	}
}

func (c *Card) GetId() string {
	return c.id
}

func (c *Card) GetName() string {
	return c.name
}

func (c *Card) GetHolderName() string {
	return c.holderName
}

func (c *Card) GetNumber() string {
	return c.number
}

func (c *Card) GetExpirationDate() time.Time {
	return c.expirationDate
}

func (c *Card) GetCvv() string {
	return c.cvv
}

func (c *Card) GetDeleted() bool {
	return c.deleted
}

func (c *Card) SetName(name string) {
	c.name = name
}

func (c *Card) SetHolderName(holderName string) {
	c.holderName = holderName
}

func (c *Card) SetNumber(number string) {
	c.number = number
}

func (c *Card) SetExpirationDate(expirationDate time.Time) {
	c.expirationDate = expirationDate
}

func (c *Card) SetDelete(deleted bool) {
	c.deleted = deleted
}
