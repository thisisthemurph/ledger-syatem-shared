package events

import "github.com/google/uuid"

type TransactionEventType string

const (
	TransactionDeposit    TransactionEventType = "transaction.deposit"
	TransactionWithdrawal TransactionEventType = "transaction.withdrawal"
)

type TransactionEvent struct {
	Type          TransactionEventType `json:"type"`
	TransactionID uuid.UUID            `json:"transactionId"`
}

type TransactionDepositEvent struct {
	TransactionEvent

	AccountID uuid.UUID `json:"accountId"`
	Amount    int64     `json:"amount"`
}

func NewTransactionDepositEvent(accountID uuid.UUID, amount int64) TransactionDepositEvent {
	return TransactionDepositEvent{
		TransactionEvent: TransactionEvent{
			Type:          TransactionDeposit,
			TransactionID: uuid.New(),
		},
		AccountID: accountID,
		Amount:    amount,
	}
}

type TransactionWithdrawalEvent struct {
	TransactionEvent

	AccountID uuid.UUID `json:"accountId"`
	Amount    int64     `json:"amount"`
}

func NewTransactionWithdrawalEvent(accountID uuid.UUID, amount int64) TransactionWithdrawalEvent {
	return TransactionWithdrawalEvent{
		TransactionEvent: TransactionEvent{
			Type:          TransactionWithdrawal,
			TransactionID: uuid.New(),
		},
		AccountID: accountID,
		Amount:    amount,
	}
}
