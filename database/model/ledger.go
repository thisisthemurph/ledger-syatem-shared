package model

import (
	"github.com/google/uuid"
	"time"
)

type TransactionStatus string

const (
	TransactionStatusPending TransactionStatus = "pending"
	TransactionStatusPosted  TransactionStatus = "posted"
	TransactionStatusFailed  TransactionStatus = "failed"
)

type Ledger struct {
	ID            uuid.UUID         `db:"id"`
	TransactionID uuid.UUID         `db:"transaction_id"`
	AccountID     uuid.UUID         `db:"account_id"`
	Amount        int64             `db:"amount"`
	Status        TransactionStatus `db:"status"`
	CreatedAt     time.Time         `db:"created_at"`
	UpdatedAt     time.Time         `db:"updated_at"`
}
