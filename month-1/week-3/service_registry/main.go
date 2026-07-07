package main

import "fmt"

type Service struct {
	ID          string
	Name        string
	Environment string
	URL         string
}

type ProbeResult struct {
	ServiceID  string
	StatusCode int
	LatencyMS  int
	Error       string
}

func isHealthy(result ProbeResult) bool {
	return result.StatusCode >= 200 && result.StatusCode < 300 && result.Error == ""
}

func printServiceStatus(service Service, result ProbeResult) {
	status := "UNHEALTHY"

	if isHealthy(result) {
		status = "HEALTHY"
	}

	fmt.Printf(
		"Service: %s | Env: %s | URL: %s | Status: %s | Code: %d | Latency: %dms\n",
		service.Name,
		service.Environment,
		service.URL,
		status,
		result.StatusCode,
		result.LatencyMS,
	)
}

func main() {
	paymentService := Service{
		ID:          "svc-001",
		Name:        "payment-api",
		Environment: "prod",
		URL:         "https://payment.example.com/health",
	}

	orderService := Service{
		ID:          "svc-002",
		Name:        "order-api",
		Environment: "staging",
		URL:         "https://order.example.com/health",
	}

	paymentResult := ProbeResult{
		ServiceID:  "svc-001",
		StatusCode: 200,
		LatencyMS:  120,
		Error:      "",
	}

	orderResult := ProbeResult{
		ServiceID:  "svc-002",
		StatusCode: 503,
		LatencyMS:  850,
		Error:      "service unavailable",
	}

	printServiceStatus(paymentService, paymentResult)
	printServiceStatus(orderService, orderResult)
}