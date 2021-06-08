package main

import (
	"flag"
	"fmt"
	"github.com/linus-capital/fineract-client/internal/fineract"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected an action like 'create'")
		os.Exit(1)
	}

	action := os.Args[1]

	create := flag.NewFlagSet("create", flag.ExitOnError)
	projectName := create.String("n", "", "name of the Project")

	update := flag.NewFlagSet("update", flag.ExitOnError)
	projectId := update.Int64("id", 0, "id of a Project")

	switch action {
	case "create":
		create.Parse(os.Args[2:])
		fineract.CreateProject(*projectName)
	case "update":
		update.Parse(os.Args[2:])
		validateProjectId(*projectId)
		fineract.UpdateProject(*projectId)
	default:
		fmt.Println("got an unexpected action!")
		os.Exit(1)
	}
}

func validateProjectId(projectId int64)  {
	if projectId == 0 {
		fmt.Println("expected an 'id' flag for the projectId")
		os.Exit(1)
	}
}