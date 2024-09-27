package model

import "time"

type Token interface {
	getExpiryDate() time.Time
}
