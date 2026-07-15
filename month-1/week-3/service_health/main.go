package main

import (
	"fmt"
)

func updateStatus(health map[string]string, service string, status string) map[string]string {
	health[service] = status
	return health
}

func getStatus(health map[string]string, service string) string {
	value, exist := health[service]
	if exist {
		return value
	} else {
		return "SERVICE_DOESNOT_EXIST"
	}
}

func countByStatus(health map[string]string) map[string]int {
	returnmap := map[string]int{}
	for _, v := range health {
		returnmap[v] = returnmap[v] + 1
	}
	return returnmap
}

func unhealthyServices(health map[string]string) []string {
	unhealthyServices := []string{}
	for k, v := range health {
		if v == "unhealthy" {
			unhealthyServices = append(unhealthyServices, k)
		}
	}
	return unhealthyServices
}

func removeService(health map[string]string, service string) {
	_, exist := health[service]
	if exist {
		delete(health, service)
	}
}

func main() {
	health := map[string]string{
		"payments-api":  "healthy",
		"users-api":     "unhealthy",
		"inventory-api": "healthy",
	}
	fmt.Println("Adding notification-api", updateStatus(health, "notification-api", "unknown"))
	fmt.Println("Checking if `search-api` exist or not: ", getStatus(health, "search-api"))
	fmt.Println("Counting each statuses: ", countByStatus(health))
	fmt.Println("List of unhealthy services ", unhealthyServices(health))
	removeService(health, "inventory-api")
	fmt.Println("After deleting `inventory-api` map is:", health)
}
