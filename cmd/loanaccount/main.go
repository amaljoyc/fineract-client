package main

import (
	"github.com/linus-capital/fineract-client/internal/fineract"
	"os"
	"strconv"
)

func main() {
	action := os.Args[1]
	var loanId int64 = 0
	if len(os.Args) > 2 {
		loanId, _ = strconv.ParseInt(os.Args[2], 10, 64)
	}

	if action == "create" {
		fineract.CreateLoan()
	} else if action == "approve" {
		validateLoanId(loanId)
		fineract.ApproveLoan(loanId)
	} else if action == "disburse" {
		validateLoanId(loanId)
		fineract.DisburseLoan(loanId)
	} else if action == "start" {
		loanId = fineract.CreateLoan()
		fineract.ApproveLoan(loanId)
		fineract.DisburseLoan(loanId)
	} else if action == "repay" {
		validateLoanId(loanId)
		fineract.RepayLoan(loanId)
	}
}

func validateLoanId(loanId int64)  {
	if loanId == 0 {
		panic("invalid loanId: 0")
	}
}
