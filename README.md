# Billing Engine

This project implements a billing system for a loan engine. The billing engine provides functionality to generate a loan schedule, track the outstanding amount for a given loan, and determine if a customer is delinquent.

## Features

- Generate a loan schedule for a specified loan.
- Track the outstanding amount for a given loan.
- Check if a customer is delinquent (more than two weeks of non-payment).
- Make payments towards the loan.

## Project Structure

billing-engine/
├── cmd/
│ └── billing-engine/
│   └─ main.go
├── pkg/
│ └── billing/
│ ├── loan.go
│ └── billing_test.go
├── go.mod
├── go.sum
└── README.md


- **cmd/**: Contains the entry point of the application.
- **pkg/**: Contains the core packages of the application. The `billing` package encapsulates all billing logic, including loan and payment-related operations.
- **go.mod**: Defines the module and its dependencies.
- **go.sum**: Contains checksums of the dependencies to ensure consistency.

## Getting Started

### Prerequisites

- Go 1.18 or higher

### Installation

1. Clone the repository:
   sh
   git clone https://github.com/fadelmajid/billing-engine.git
   cd billing-engine
2. Initialize the module:
   go mod tidy


## Usage
1. To run the application:
    run cmd/billing-engine/main.go

2. To run the test:
    go test ./pkg/billing

## API
# Loan
NewLoan(loanID int, principal float64, interestRate float64, durationWeeks int) *Loan
Creates a new loan with the specified parameters.

# GetOutstanding() float64
Returns the current outstanding amount on the loan.

# IsDelinquent() bool
Checks if the borrower is delinquent (missed more than two consecutive payments).

# MakePayment(amount float64) string
Processes a payment for the loan. Returns a message indicating whether the payment was successful or not.


## Example
package main

import (
    "fmt"
    "github.com/yourusername/billing-engine/pkg/billing"
)

func main() {
    loan := billing.NewLoan(100, 5000000, 0.1, 50)
    fmt.Println(loan.MakePayment(110000)) // Payment successful
    fmt.Println(loan.GetOutstanding())    // 5390000
    fmt.Println(loan.IsDelinquent())      // False
}