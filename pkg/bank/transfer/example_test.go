package transfer

import "fmt"

func ExampleTotal() {
	fmt.Println(Total(0))
	fmt.Println(Total(5_000))
	fmt.Println(Total(10_000))

	// Output:
	// 0
	// 5025
	// 10050
}
