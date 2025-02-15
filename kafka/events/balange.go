package events

import "github.com/google/uuid"

type BalanceUpdatedEvent struct {
	AccountID uuid.UUID `json:"accountId"`
	Balance   int64     `json:"balance"`
}

func NewBalanceUpdatedEvent(accountID uuid.UUID, balance int64) BalanceUpdatedEvent {
	return BalanceUpdatedEvent{
		AccountID: accountID,
		Balance:   balance,
	}
}
