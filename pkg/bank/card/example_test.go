package card

import (
	"bank/pkg/types"
	"fmt"
)

func ExampleWithdraw() {
	card := types.Card{Balance: 30_000_00, Active: true}
	Withdraw(&card, 5_000_00)

	fmt.Println(card.Balance)
	// Output: 2500000
}

func ExampleWithdraw_positive() {
	card := types.Card{Balance: 30_000_00, Active: true}
	Withdraw(&card, -3_000_00)

	fmt.Println(card.Balance)
	// Output: 3000000
}

func ExampleWithdraw_noMoney() {
	card := types.Card{Balance: 10_000_00, Active: true}
	Withdraw(&card, 15_000_00)

	fmt.Println(card.Balance)
	// Output: 1000000
}

func ExampleWithdraw_inactive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	card.Active = false
	Withdraw(&card, 3_000_00)

	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleWithdraw_limit() {
	card := types.Card{Balance: 30_000_00, Active: true}
	Withdraw(&card, 25_000_00)

	fmt.Println(card.Balance)
	// Output: 3000000
}

func ExampleDeposit() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 5_000_00)

	fmt.Println(card.Balance)
	// Output: 2500000
}

func ExampleDeposit_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Deposit(&card, 10_000_00)

	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleDeposit_limit() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 60_000_00)

	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleAddBonus() {
	card := types.Card{Balance: 20_000_00, Active: true, MinBalance: 10_000_00}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: 2002465
}

func ExampleAddBonus_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false, MinBalance: 10_000_00}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleAddBonus_negativeBalance() {
	card := types.Card{Balance: -20_000_00, Active: true, MinBalance: 10_000_00}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: -2000000
}

func ExampleAddBonus_limit() {
	card := types.Card{Balance: 15_000_00, Active: true, MinBalance: 10_000_00}
	AddBonus(&card, 1000, 30, 365)

	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleTotal() {
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
	}))

	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
		{
			Balance: 2_000_00,
			Active:  true,
		},
	}))

	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  false,
		},
	}))

	fmt.Println(Total([]types.Card{
		{
			Balance: -1_000_00,
			Active:  true,
		},
	}))

	// Output:
	// 100000
	// 300000
	// 0
	// 0
}

func ExampleTotal_negative() {
	cards := []types.Card{
		{
			Balance: -10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
	}

	fmt.Println(Total(cards))
	// Output: 1000000
}

func ExampleTotal_inactive() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  false,
		},
	}
	fmt.Println(Total(cards))
	// Output: 1000000
}
