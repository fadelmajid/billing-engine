package billing

import "time"

// Payment represents a single payment
type Payment struct {
	Week       int
	AmountDue  float64
	AmountPaid float64
	DueDate    time.Time
}

// Loan represents a loan with multiple payments
type Loan struct {
	LoanID        int
	Principal     float64
	InterestRate  float64
	DurationWeeks int
	TotalAmount   float64
	WeeklyPayment float64
	Payments      []Payment
	CurrentWeek   int
}

// NewLoan initializes a new Loan
func NewLoan(loanID int, principal float64, interestRate float64, durationWeeks int) *Loan {
	totalAmount := principal * (1 + interestRate)
	weeklyPayment := totalAmount / float64(durationWeeks)
	payments := make([]Payment, durationWeeks)

	for i := 0; i < durationWeeks; i++ {
		payments[i] = Payment{
			Week:      i + 1,
			AmountDue: weeklyPayment,
			DueDate:   time.Now().AddDate(0, 0, 7*(i+1)),
		}
	}

	return &Loan{
		LoanID:        loanID,
		Principal:     principal,
		InterestRate:  interestRate,
		DurationWeeks: durationWeeks,
		TotalAmount:   totalAmount,
		WeeklyPayment: weeklyPayment,
		Payments:      payments,
		CurrentWeek:   1,
	}
}

// GetOutstanding returns the current outstanding amount on the loan
func (loan *Loan) GetOutstanding() float64 {
	outstanding := 0.0
	for _, payment := range loan.Payments {
		outstanding += payment.AmountDue - payment.AmountPaid
	}
	return outstanding
}

// IsDelinquent checks if the borrower is delinquent
func (loan *Loan) IsDelinquent() bool {
	consecutiveMissedPayments := 0
	for _, payment := range loan.Payments {
		if payment.AmountPaid < payment.AmountDue {
			consecutiveMissedPayments++
			if consecutiveMissedPayments >= 2 {
				return true
			}
		} else {
			consecutiveMissedPayments = 0
		}
	}
	return false
}

// MakePayment processes a payment for the loan
func (loan *Loan) MakePayment(amount float64) string {
	for i := range loan.Payments {
		if loan.Payments[i].AmountPaid < loan.Payments[i].AmountDue {
			if amount == loan.Payments[i].AmountDue-loan.Payments[i].AmountPaid {
				loan.Payments[i].AmountPaid += amount
				return "Payment successful"
			} else {
				return "Incorrect payment amount, must pay the exact due amount for the week"
			}
		}
	}
	return "No outstanding payments"
}
