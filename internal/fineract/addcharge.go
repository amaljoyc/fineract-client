package fineract

import (
	"encoding/json"
	"fmt"
	fineractclient "github.com/linus-capital/fineract-client"
	"github.com/linus-capital/fineract-client/internal/util"
	"strconv"
)

const addChargeUrl = fineractclient.FineractApi + "/loans/%s/charges"

type addChargeResponse struct {
	ResourceId int64
}

func AddCharge(startDate string, dueDate string, chargeAmount float64, allowRecalculation bool, chargeLoanId int64)  {
	dataMap := readAddChargeTemplate()
	if startDate != "" {
		dataMap["startDate"] = startDate
	}
	if dueDate != "" {
		dataMap["dueDate"] = dueDate
	}
	if chargeAmount != 0 {
		dataMap["amount"] = chargeAmount
	}

	dataMap["allowRecalculation"] = allowRecalculation
	response := addCharge(dataMap, chargeLoanId)
	fmt.Println("Added new fee/charge with id", response.ResourceId)
}

func readAddChargeTemplate() map[string]interface{} {
	data := util.Read("addCharge.json")
	var dataMap map[string]interface{}
	json.Unmarshal(data, &dataMap)
	return dataMap
}

func addCharge(dataMap map[string]interface{}, loanId int64) addChargeResponse  {
	body, _ := json.Marshal(dataMap)
	var response addChargeResponse
	url := fmt.Sprintf(addChargeUrl, strconv.FormatInt(loanId, 10))
	responseJson := util.Request("POST", url, body)
	err := json.Unmarshal(responseJson, &response)
	if err != nil {
		panic(err)
	}
	return response
}