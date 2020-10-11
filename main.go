package main

import (
	"fmt"
	"time"
	"github.com/kpfaulkner/azurecosts/pkg"
)

func main() {
	fmt.Printf("so it begins...\n")

	subscriptionID := ""
	tenantID := ""
	clientID := ""
	clientSecret := ""

	ac := pkg.NewAzureCost(subscriptionID, tenantID, clientID, clientSecret)

	startDate := time.Date(2020, 04, 01, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2020, 04, 30, 0, 0, 0, 0, time.UTC)

	subscriptionCosts, err := ac.GenerateSubscriptionCostDetails([]string{subscriptionID}, startDate, endDate)
	if err != nil {
		fmt.Printf("ERROR %s\n", err.Error())
	}

	total := 0.0
	for _, sc := range subscriptionCosts {
		for k, v := range sc.ResourceGroupCosts {
			fmt.Printf("RG %s cost %0.2f\n", k, v)
		}
		total += sc.Total
		fmt.Printf("Sub %s cost %0.2f\n", sc.SubscriptionID, sc.Total)
	}
	fmt.Printf("TOTAL for everyting was %0.2f\n", total)

	azureCostsData := ac.FilterDataBasedOnSubscription(subscriptionCosts, []string{subscriptionID})

	// filter costs.... just an example
	prefixCosts, _ := ac.GetCostsPerRGPrefix([]string{"test-"}, azureCostsData)
	for k, v := range prefixCosts {
		fmt.Printf("Prefix %s has cost %0.2f\n", k, v)
	}

	fmt.Printf("azure total costs %0.2f\n", subscriptionCosts[0].Total)
}
