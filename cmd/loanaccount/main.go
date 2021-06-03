package main

import (
	"flag"
	"fmt"
	"github.com/linus-capital/fineract-client/internal/fineract"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected an action like 'create' or 'open'")
		os.Exit(1)
	}

	action := os.Args[1]
	repay := flag.NewFlagSet("repay", flag.ExitOnError)
	loanId := repay.Int64("id", 0, "loanId of a loanAccount")
	principal := repay.Float64("p", 0, "value of principalPortionGiven")
	interest := repay.Float64("i", 0, "value of interestPortionGiven")
	fee := repay.Float64("f", 0, "value of feePortionGiven")
	amount := repay.Float64("a", 0, "value of transactionAmount")
	date := repay.String("d", "", "value of transactionDate")

	switch action {
	case "create":
		fineract.CreateLoan()
	case "open":
		loanId := fineract.CreateLoan()
		fineract.ApproveLoan(loanId)
		fineract.DisburseLoan(loanId)
	case "repay":
		repay.Parse(os.Args[2:])
		validateLoanId(*loanId)
		fineract.RepayLoan(*loanId, *principal, *interest, *fee, *amount, *date)
	default:
		fmt.Println("got an unexpected action!")
		os.Exit(1)
	}
}

func validateLoanId(loanId int64)  {
	if loanId == 0 {
		fmt.Println("expected an 'id' flag for the loanId")
		os.Exit(1)
	}
}
