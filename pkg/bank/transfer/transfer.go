package transfer

import "bank/pkg/types"

const bonusPercent = 0.0050

// Bonus calculates bonus by transfer amount
func Bonus(amount types.Money) types.Money {
	bonus := types.Money(float64(amount) * bonusPercent)
	return bonus
}

// Total calculates total amount by transfer taking into bonus
func Total(amount types.Money) types.Money {
	total := amount + Bonus(amount)
	return total
}
