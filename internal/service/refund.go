package service

import (
	"banking-api/internal/entity"
	"errors"
)

var (
	TransactionAlreadyRefunded = errors.New("transaction already refunded")
	TransactionNotFound        = errors.New("transaction not found")
)

type RefundService interface {
	Create(origin *entity.RefundRequest) (*entity.RefundResponse, error)
}

func NewRefundService() *refundService {
	return &refundService{}
}

type refundService struct {
}

func (s *refundService) Create(req *entity.RefundRequest) (*entity.RefundResponse, error) {
	if err := validateTransactionNotFound(req.TransactionID); err != nil {
		return nil, err
	}
	if err := validateTransactionAlreadyRefunded(req.TransactionID); err != nil {
		return nil, err
	}
	return entity.MapRefundRequestToResponse(req), nil
}

func validateTransactionAlreadyRefunded(TransactionID string) error {
	if TransactionID == "1234-5678-9012-3456" {
		return TransactionAlreadyRefunded
	}
	return nil
}
func validateTransactionNotFound(TransactionID string) error {
	if TransactionID == "1234-5678-9012-3457" {
		return TransactionNotFound
	}
	return nil
}
