package entity

import (
	"banking-api/pkg/util"
	"time"
)

func MapPaymentRequestToResponse(req *PaymentRequest) *PaymentResponse {
	return &PaymentResponse{
		TransactionID: api_util.GenerateUUID(),
		Amount:        req.Amount,
		Currency:      req.Currency,
		Merchant:      req.Merchant,
		Timestamp:     time.Now(),
	}
}

type PaymentResponse struct {
	TransactionID string    `json:"transactionId"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Merchant      string    `json:"merchant"`
	Timestamp     time.Time `json:"timestamp"`
}

type PaymentRequest struct {
	CardID      string  `json:"cardId"`
	CVC         string  `json:"cvc"`
	ExpiredDate string  `json:"expiredDate"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Merchant    string  `json:"merchant"`
}
