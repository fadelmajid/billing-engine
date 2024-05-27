package main

import (
	"fmt"
	"github.com/fadelmajid/billing-engine/pkg/billing"
)

func main() {
	loan := billing.NewLoan(100, 5000000, 0.1, 50)

	// Making a payment of 110000 for the first week
	fmt.Println(loan.MakePayment(110000)) // Output: Payment successful

	// Checking outstanding amount after one payment
	fmt.Println(loan.GetOutstanding()) // Output: 5390000

	// Making another payment of 110000 for the second week
	fmt.Println(loan.MakePayment(110000)) // Output: Payment successful

	// Checking if the borrower is delinquent (should be False)
	fmt.Println(loan.IsDelinquent()) // Output: False

	// Simulating missed payments and checking delinquency
	for i := 0; i < 2; i++ { // Skipping two weeks of payment
		loan.CurrentWeek++
	}

	fmt.Println(loan.IsDelinquent()) // Output: True
}
