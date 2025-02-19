package card

import (
	"bank/pkg/types"
)

const withdrawLimit = 20_000_00
const depositLimit = 50_000_00
const bonusLimit = 5_000_00

func Withdraw(card *types.Card, amount types.Money) {
	if !card.Active {
		return
	}

	if amount < 0 {
		return
	}

	if amount > withdrawLimit {
		return
	}

	if card.Balance < amount {
		return
	}

	card.Balance -= amount

	return
}

func Deposit(card *types.Card, amount types.Money) {
	if !card.Active {
		return
	}

	if amount > depositLimit {
		return
	}

	card.Balance += amount

	return
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !card.Active {
		return
	}

	if card.Balance <= 0 {
		return
	}
	bonus := (int(card.MinBalance) * percent * daysInMonth) / (100 * daysInYear)

	if bonus > bonusLimit {
		bonus = bonusLimit
	}

	card.Balance += types.Money(bonus)

	return
}

// Total вычисляет общую сумму средств на всех картах.
// Отрицательные суммы и суммы на заблокированных картах игнорируюся.
func Total(cards []types.Card) types.Money {
	sum := types.Money(0)
	for _, card := range cards {
		if card.Balance < 0 {
			continue
		}
		if !card.Active {
			continue
		}
		sum += card.Balance
	}
	return sum
}

func PaymentSource(cards []types.Card) []types.PaymentSource {
	var sources []types.PaymentSource
	for _, card := range cards {
		if card.Balance <= 0 {
			continue
		}

		if !card.Active {
			continue
		}

		source := types.PaymentSource{Type: "card", Number: string(card.PAN), Balance: card.Balance}

		sources = append(sources, source)
	}

	return sources
}
