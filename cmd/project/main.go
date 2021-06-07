package main

import (
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

	switch action {
	case "create":
		fineract.CreateProject()
	default:
		fmt.Println("got an unexpected action!")
		os.Exit(1)
	}
}