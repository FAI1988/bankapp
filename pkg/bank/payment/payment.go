package payment

import (
	"bank/pkg/types"
)

// Max возвращает максимальный платёж
func Max(payments []types.Payment) types.Payment {
	maxAmount := types.Money(0)

	for _, payment := range payments {
		if payment.Amount > maxAmount {
			maxAmount = payment.Amount
		}
	}

	for _, payment := range payments {
		if payment.Amount == maxAmount {
			return payment
		}
	}
	return types.Payment{}
}
