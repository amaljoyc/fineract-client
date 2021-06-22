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

	charge := flag.NewFlagSet("charge", flag.ExitOnError)
	chargeLoanId := charge.Int64("id", 0, "loanId of a loanAccount to add charge to")
	startDate := charge.String("s", "", "startDate of charge")
	dueDate := charge.String("d", "", "dueDate of charge")
	chargeAmount := charge.Float64("a", 0, "amount of charge")
	allowRecalculation := charge.Bool("r", false, "allow recalculation or not (true|false)")

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
	case "charge":
		charge.Parse(os.Args[2:])
		fineract.AddCharge(*startDate, *dueDate, *chargeAmount, *allowRecalculation, *chargeLoanId)
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
