package fineract

import (
	"encoding/json"
	"fmt"
	fineractclient "github.com/linus-capital/fineract-client"
	"github.com/linus-capital/fineract-client/internal/util"
)

type CreateProjectResponse struct {
	ClientId int64
}

const projectUrl = fineractclient.FineractApi + "/clients"

func CreateProject()  {
	data := util.Read("createProject.json")
	var response CreateProjectResponse
	responseJson := util.Request(projectUrl, data)
	err := json.Unmarshal(responseJson, &response)
	if err != nil {
		panic(err)
	}

	fmt.Println("Created new project with id", response.ClientId)
}
