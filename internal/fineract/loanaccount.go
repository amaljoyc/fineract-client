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

type RepayLoanRequest struct {
	PaymentTypeId         int    `json:"paymentTypeId"`
	TransactionAmount     float64    `json:"transactionAmount,omitempty"`
	PrincipalPortionGiven float64    `json:"principalPortionGiven,omitempty"`
	InterestPortionGiven  float64    `json:"interestPortionGiven,omitempty"`
	FeePortionGiven       float64    `json:"feePortionGiven,omitempty"`
	TransactionDate       string `json:"transactionDate"`
	Locale                string `json:"locale"`
	DateFormat            string `json:"dateFormat"`
}

const loanUrl = fineractclient.FineractApi + "/loans/"

func CreateLoan() int64  {
	data := util.Read("createLoan.json")
	var response CreateLoanResponse
	responseJson := util.Request("POST", loanUrl, data)
	err := json.Unmarshal(responseJson, &response)
	if err != nil {
		 panic(err)
	}

	fmt.Println("Created new loan with id", response.LoanId)
	return response.LoanId
}

func ApproveLoan(loanId int64)  {
	data := util.Read("approveLoan.json")
	util.Request("POST",loanUrl + strconv.FormatInt(loanId, 10) + "?command=approve", data)
	fmt.Println("Approved loan with id", loanId)
}

func DisburseLoan(loanId int64)  {
	data := util.Read("disburseLoan.json")
	util.Request("POST",loanUrl + strconv.FormatInt(loanId, 10) + "?command=disburse", data)
	fmt.Println("Disbursed loan with id", loanId)
}

func RepayLoan(loanId int64, principal float64, interest float64, fee float64, amount float64, date string)  {
	data := util.Read("repayLoan.json")
	var repayLoanRequest RepayLoanRequest
	err := json.Unmarshal(data, &repayLoanRequest)
	if err != nil {
		panic(err)
	}

	if principal != 0 {
		repayLoanRequest.PrincipalPortionGiven = principal
	}
	if interest != 0 {
		repayLoanRequest.InterestPortionGiven = interest
	}
	if fee != 0 {
		repayLoanRequest.FeePortionGiven = fee
	}
	if amount != 0 {
		repayLoanRequest.TransactionAmount = amount
	}
	if date != "" {
		repayLoanRequest.TransactionDate = date
	}

	body, _ := json.Marshal(repayLoanRequest)
	util.Request("POST",loanUrl + strconv.FormatInt(loanId, 10) + "/transactions?command=repayment", body)
	fmt.Println("Repaid loan with id", loanId)
}
