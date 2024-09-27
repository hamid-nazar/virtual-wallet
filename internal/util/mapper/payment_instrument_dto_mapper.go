package util

import (
	"github.com/hamid-nazari/virtual-wallet/internal/dto"
	"github.com/hamid-nazari/virtual-wallet/internal/model"
)

func MapToPaymentInstrumentDto(c dto.NewCardDto) model.Card {
	return NewCard(c.Number, c.CardHolder, c.ExpirationDate, c.CVV)
}
