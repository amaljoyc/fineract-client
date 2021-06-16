package fineract

import (
	"encoding/json"
	"fmt"
	fineractclient "github.com/linus-capital/fineract-client"
	"github.com/linus-capital/fineract-client/internal/util"
)

const chargeUrl = fineractclient.FineractApi + "/charges"

type chargeResponse struct {
	ResourceId int64
}

func CreateFlatFee(amount float64)  {
	dataMap := readChargeTemplate()
	dataMap["name"] = "FlatFee on DueDate"
	dataMap["amount"] = amount
	dataMap["chargeCalculationType"] = 1 // flat
	dataMap["linusStyleCharge"] = false

	response := createCharge(dataMap)
	fmt.Println("Created new FLAT fee with id", response.ResourceId)
}

func CreatePercentFee(amount float64)  {
	dataMap := readChargeTemplate()
	dataMap["name"] = "PercentFee on DueDate"
	dataMap["amount"] = amount
	dataMap["chargeCalculationType"] = 2 // amountPercent
	dataMap["linusStyleCharge"] = true

	response := createCharge(dataMap)
	fmt.Println("Created new PERCENT fee with id", response.ResourceId)
}

func readChargeTemplate() map[string]interface{} {
	data := util.Read("createCharge.json")
	var dataMap map[string]interface{}
	json.Unmarshal(data, &dataMap)
	return dataMap
}

func createCharge(dataMap map[string]interface{}) chargeResponse  {
	body, _ := json.Marshal(dataMap)
	var response chargeResponse
	responseJson := util.Request("POST", chargeUrl, body)
	err := json.Unmarshal(responseJson, &response)
	if err != nil {
		panic(err)
	}
	return response
}