package service

import (
	"github.com/hamid-nazari/virtual-wallet/internal/dto"
)

type CardService struct {
	userService              *UserService
	paymentInstrumentService *PaymentInstrumentService
}

func NewCardService(UserService *UserService, PaymentInstrumentService *PaymentInstrumentService) *CardService {
	return &CardService{
		userService:              UserService,
		paymentInstrumentService: PaymentInstrumentService,
	}

}

func (c *CardService) Create(cardHolderUsername string, c *dto.NewCardDto) error {
	user := c.UserService.GetByUsername(cardHolderUsername)

}
