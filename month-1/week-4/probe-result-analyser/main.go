package main

import (
	"fmt"
)

type ProbeResult struct {
	Endpoint       string `json:"endpoint"`
	HttpMethod     string `json:"httpmethod"`
	TimeLatency    int    `json:"latency"`
	HttpStatusCode int    `json:"statusCode"`
	AttemptNumber  int    `json:"attempt"`
}

func (pr *ProbeResult) UpdateProbeResult() {
	pr.Endpoint = "http://api.payments.com"
	pr.HttpMethod = "POST"
	pr.HttpStatusCode = 200
	pr.AttemptNumber = 1
	pr.TimeLatency = 23
}

func main() {
	probeRs := ProbeResult{}
	probeRs.UpdateProbeResult()
	fmt.Println("Probe Data: ", probeRs)
}
