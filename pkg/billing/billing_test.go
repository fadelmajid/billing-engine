package billing

import (
	"testing"
)

func TestLoan_GetOutstanding(t *testing.T) {
	loan := NewLoan(100, 5000000, 0.1, 50)

	if loan.GetOutstanding() != 5500000 {
		t.Errorf("Expected outstanding amount to be 5500000, got %f", loan.GetOutstanding())
	}
}

func TestLoan_IsDelinquent(t *testing.T) {
	loan := NewLoan(100, 5000000, 0.1, 50)
	loan.MakePayment(110000)
	loan.MakePayment(110000)

	if loan.IsDelinquent() {
		t.Errorf("Expected borrower to not be delinquent")
	}

	// Simulating missed payments
	loan.CurrentWeek += 2

	if !loan.IsDelinquent() {
		t.Errorf("Expected borrower to be delinquent")
	}
}

func TestLoan_MakePayment(t *testing.T) {
	loan := NewLoan(100, 5000000, 0.1, 50)
	result := loan.MakePayment(110000)

	if result != "Payment successful" {
		t.Errorf("Expected 'Payment successful', got %s", result)
	}

	if loan.GetOutstanding() != 5390000 {
		t.Errorf("Expected outstanding amount to be 5390000, got %f", loan.GetOutstanding())
	}
}
