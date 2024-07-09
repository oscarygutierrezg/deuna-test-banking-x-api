package entity

import (
	"banking-api/pkg/util"
	"time"
)

func MapRefundRequestToResponse(req *RefundRequest) *RefundResponse {
	return &RefundResponse{
		RefundID:      api_util.GenerateUUID(),
		TransactionID: req.TransactionID,
		Amount:        req.Amount,
		Currency:      req.Currency,
		Timestamp:     time.Now(),
	}
}

type RefundResponse struct {
	TransactionID string    `json:"transactionId"`
	RefundID      string    `json:"refundId"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Timestamp     time.Time `json:"timestamp"`
}

type RefundRequest struct {
	TransactionID string  `json:"transactionId"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
}
