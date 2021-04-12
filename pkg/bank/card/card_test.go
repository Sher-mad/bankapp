package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	card := []types.Card{
		{
			PAN:     "5058 xxxx xxxx 9809",
			Balance: 1500,
			Active:  false,
		},
		{
			PAN:     "5058 xxxx xxxx 9988",
			Balance: 1200,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 9880",
			Balance: 1900,
			Active:  true,
		},
	}

	payments := PaymentSources(card)

	for _, payment := range payments {
		fmt.Println(payment.Number)

	}

	// Output:
	// 5058 xxxx xxxx 9988
	// 5058 xxxx xxxx 9880

}
