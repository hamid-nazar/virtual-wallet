package service

import (
	"github.com/hamid-nazari/virtual-wallet/internal/model"
)

type UserService struct {
	WalletService *WalletService
}

func NewUserService(WalletService *WalletService) *UserService {
	return &UserService{
		WalletService: WalletService,
	}
}

func (u *UserService) Create(username string) error {
	return nil
}
func (u *UserService) GetByUsername(username string) model.User {
	return nil
}
