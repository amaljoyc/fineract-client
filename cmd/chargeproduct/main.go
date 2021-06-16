package main

import (
	"flag"
	"fmt"
	"github.com/linus-capital/fineract-client/internal/fineract"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected an action like 'flat' or 'percent'")
		os.Exit(1)
	}

	action := os.Args[1]

	flat := flag.NewFlagSet("flat", flag.ExitOnError)
	flatAmount := flat.Float64("a", 0, "flat amount")

	percent := flag.NewFlagSet("percent", flag.ExitOnError)
	percentAmount := percent.Float64("a", 0, "percent amount")

	switch action {
	case "flat":
		flat.Parse(os.Args[2:])
		validateAmount(*flatAmount)
		fineract.CreateFlatFee(*flatAmount)
	case "percent":
		percent.Parse(os.Args[2:])
		validateAmount(*percentAmount)
		fineract.CreatePercentFee(*percentAmount)
	default:
		fmt.Println("got an unexpected action!")
		os.Exit(1)
	}
}

func validateAmount(amount float64)  {
	if amount == 0 {
		fmt.Println("expected an 'a' flag for the amount")
		os.Exit(1)
	}
}