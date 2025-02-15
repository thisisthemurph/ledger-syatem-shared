package events

import (
	"errors"
	"github.com/google/uuid"
)

var ErrInvalidTransactionAmount = errors.New("invalid transaction amount")

type TransactionEventType string

const (
	TransactionDeposit    TransactionEventType = "transaction.deposit"
	TransactionWithdrawal TransactionEventType = "transaction.withdrawal"
	TransactionTransfer   TransactionEventType = "transaction.transfer"
)

type TransactionEvent interface {
	Type() TransactionEventType
	TransactionID() uuid.UUID
	AccountID() uuid.UUID
	Amount() int64
}

func NewDepositTransactionEvent(accountID uuid.UUID, amount int64) TransactionEvent {
	return SingeTransactionEvent{
		TransactionType:       TransactionDeposit,
		TransactionIdentifier: uuid.New(),
		AccountIdentifier:     accountID,
		TransactionAmount:     amount,
	}
}

func NewWithdrawalTransactionEvent(accountID uuid.UUID, amount int64) TransactionEvent {
	return SingeTransactionEvent{
		TransactionType:       TransactionWithdrawal,
		TransactionIdentifier: uuid.New(),
		AccountIdentifier:     accountID,
		TransactionAmount:     amount,
	}
}

func NewTransferTransactionEvent(fromAccountID, toAccountID uuid.UUID, amount int64) TransactionEvent {
	return TransferTransactionEvent{
		TransactionIdentifier: uuid.New(),
		FromAccountID:         fromAccountID,
		ToAccountID:           toAccountID,
		TransactionAmount:     amount,
	}
}

type SingeTransactionEvent struct {
	TransactionType       TransactionEventType `json:"type"`
	TransactionIdentifier uuid.UUID            `json:"transactionId"`
	AccountIdentifier     uuid.UUID            `json:"accountId"`
	TransactionAmount     int64                `json:"amount"`
}

func (s SingeTransactionEvent) Type() TransactionEventType {
	return s.TransactionType
}

func (s SingeTransactionEvent) TransactionID() uuid.UUID {
	return s.TransactionIdentifier
}

func (s SingeTransactionEvent) AccountID() uuid.UUID {
	return s.AccountIdentifier
}

func (s SingeTransactionEvent) Amount() int64 {
	return s.TransactionAmount
}

func (s SingeTransactionEvent) Validate() error {
	if s.TransactionAmount <= 0 {
		return ErrInvalidTransactionAmount
	}

	return nil
}

type TransferTransactionEvent struct {
	TransactionIdentifier uuid.UUID `json:"transactionId"`
	FromAccountID         uuid.UUID `json:"fromAccountId"`
	ToAccountID           uuid.UUID `json:"toAccountId"`
	TransactionAmount     int64     `json:"amount"`
}

func (t TransferTransactionEvent) Type() TransactionEventType {
	return TransactionTransfer
}

func (t TransferTransactionEvent) TransactionID() uuid.UUID {
	return t.TransactionIdentifier
}

func (t TransferTransactionEvent) AccountID() uuid.UUID {
	return t.FromAccountID
}

func (t TransferTransactionEvent) Amount() int64 {
	return t.TransactionAmount
}
