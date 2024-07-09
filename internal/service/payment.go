package service

import (
	"banking-api/internal/entity"
	"errors"
)

var (
	InvalidCardDetails = errors.New("invalid card details")
	InsufficientFunds  = errors.New("insufficient funds")
)

type PaymentService interface {
	Create(origin *entity.PaymentRequest) (*entity.PaymentResponse, error)
}

func NewPaymentService() *paymentService {
	return &paymentService{}
}

type paymentService struct {
}

func (s *paymentService) Create(req *entity.PaymentRequest) (*entity.PaymentResponse, error) {
	if err := validateCardDetails(req.CardID); err != nil {
		return nil, err
	}
	if err := checkCardBalance(req.Amount); err != nil {
		return nil, err
	}
	return entity.MapPaymentRequestToResponse(req), nil
}

func validateCardDetails(cardID string) error {
	if cardID == "1234-5678-9012-3456" {
		return InvalidCardDetails
	}
	return nil
}

func checkCardBalance(amount float64) error {
	if amount > 1000.00 {
		return InsufficientFunds
	}
	return nil
}
