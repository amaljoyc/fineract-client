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
		fineract.RepayLoan(*loanId)
	default:
		fmt.Println("got an unexpected action!")
		os.Exit(1)
	}
}

func validateLoanId(loanId int64)  {
	if loanId == 0 {
		fmt.Println("repay action expects a loanId")
		os.Exit(1)
	}
}
