package fineract

import (
	"encoding/json"
	"fmt"
	fineractclient "github.com/linus-capital/fineract-client"
	"github.com/linus-capital/fineract-client/internal/util"
	"strconv"
)

type CreateLoanResponse struct {
	LoanId int64
}

const loanUrl = fineractclient.FineractApi + "/loans/"

func CreateLoan() int64  {
	data := util.Read("createLoanAccount.json")
	var response CreateLoanResponse
	responseJson := util.Request(loanUrl, data)
	err := json.Unmarshal(responseJson, &response)
	if err != nil {
		 panic(err)
	}

	fmt.Println("Created new loan with id", response.LoanId)
	return response.LoanId
}

func ApproveLoan(loanId int64)  {
	data := util.Read("approveLoanAccount.json")
	util.Request(loanUrl + strconv.FormatInt(loanId, 10) + "?command=approve", data)
	fmt.Println("Approved loan with id", loanId)
}

func DisburseLoan(loanId int64)  {
	data := util.Read("disburseLoanAccount.json")
	util.Request(loanUrl + strconv.FormatInt(loanId, 10) + "?command=disburse", data)
	fmt.Println("Disbursed loan with id", loanId)
}

func RepayLoan(loanId int64)  {
	fmt.Println("RepayLoan TODO", loanId)
}
