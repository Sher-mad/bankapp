package card

import "bank/pkg/bank/types"

func PaymentSources(cards []types.Card) []types.PaymentSource {
	payments := []types.PaymentSource{}

	for _, cards := range cards {
		if !cards.Active {
			continue
		}
		if cards.Balance <= 0 {
			continue
		}
		payment := types.PaymentSource{
			Type:    "card",
			Number:  cards.PAN,
			Balance: int(cards.Balance),
		}

		payments = append(payments, payment)

	}

	return payments
}
