package card

import "bank/pkg/bank/types"

func PaymentSources(cards []types.Card) []types.PaymentSource {
	payments :=[]types.PaymentSource{}

	for _, cards := range cards {
		if  cards.Balance <=0 {
			return payments
		}
		if !cards.Active {
			return payments
		}
		payment :=types.PaymentSource{
			Type: "card",
			Number: cards.PAN,
			Balance: int(cards.Balance),
		}
		payments = append(payments, payment)
		
	}

	return payments
}
