package fineract

import (
	"encoding/json"
	"fmt"
	"github.com/linus-capital/fineract-client/internal/util"
)

type CreateLoanResponse struct {
	LoanId int64
}

func CreateLoan() int64  {
	data := util.Read("createLoanAccount.json")
	var response CreateLoanResponse
	responseJson := util.Request("https://localhost:8443/fineract-provider/api/v1/loans", data)
	err := json.Unmarshal(responseJson, &response)
	if err != nil {
		 panic(err)
	}

	fmt.Println("Created new loan with id", response.LoanId)
	return response.LoanId
}

func ApproveLoan(loanId int64)  {
	fmt.Println("ApproveLoan TODO", loanId)
}

func DisburseLoan(loanId int64)  {
	fmt.Println("DisburseLoan TODO", loanId)
}

func RepayLoan(loanId int64)  {
	fmt.Println("RepayLoan TODO", loanId)
}
