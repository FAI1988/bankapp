package payment

import (
	"bank/pkg/types"
	"fmt"
)

func ExampleMax() {
	//
	payments := []types.Payment{{1, 1000}, {99, 10000}, {3, 4000}}
	maximum := Max(payments)
	fmt.Println(maximum.ID)
	// Output: 99
}
