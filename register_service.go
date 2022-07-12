package main

import (
	"github.com/hashicorp/consul/api"
)

func main() {
	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	//Get agent
	agent := client.Agent()
	
	//Create check
	check := &api.AgentServiceCheck{
		Args: []string{"go", "run", "google_health_check.go"},
		Interval: "5s",
        Timeout: "1s",
	}

	//Register service
	reg := &api.AgentServiceRegistration{
		Name: "google",
		Port: 80,
		Check: check,
	}
	if err := agent.ServiceRegister(reg); err != nil {
		panic(err)
	}
}
