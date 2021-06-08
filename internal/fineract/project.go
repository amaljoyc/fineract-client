package fineract

import (
	"encoding/json"
	"fmt"
	fineractclient "github.com/linus-capital/fineract-client"
	"github.com/linus-capital/fineract-client/internal/util"
	"strconv"
)

type ProjectResponse struct {
	ClientId int64
}

const projectUrl = fineractclient.FineractApi + "/clients"

func CreateProject(projectName string)  {
	data := util.Read("project.json")
	var dataMap map[string]interface{}
	json.Unmarshal(data, &dataMap)
	if projectName != "" {
		dataMap["fullname"] = projectName
	}

	body, _ := json.Marshal(dataMap)
	var response ProjectResponse
	responseJson := util.Request("POST", projectUrl, body)
	err := json.Unmarshal(responseJson, &response)
	if err != nil {
		panic(err)
	}

	fmt.Println("Created new project with id", response.ClientId)
}

func UpdateProject(projectId int64)  {
	data := util.Read("project.json")
	var dataMap map[string]interface{}
	json.Unmarshal(data, &dataMap)
	delete(dataMap, "officeId")

	body, _ := json.Marshal(dataMap)
	var response ProjectResponse
	responseJson := util.Request("PUT", projectUrl + "/" + strconv.FormatInt(projectId, 10), body)
	err := json.Unmarshal(responseJson, &response)
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated project with id", response.ClientId)
}
