package util

import (
	"github.com/hashicorp/consul/api"
	"log"
)

var ConsulClient *api.Client

func init() {
	config := api.DefaultConfig()

	config.Address = "172.18.35.60:8500"

	client, error := api.NewClient(config)
	if error != nil {
		log.Fatal(error)
	}

	ConsulClient = client
}

func RegService() {
	reg := api.AgentServiceRegistration{}
	reg.ID = "ConsulService"
	reg.Name = "TestService"
	reg.Address = "172.18.32.111"
	reg.Port = 8090
	reg.Tags = []string{"primary"}

	check := api.AgentServiceCheck{}
	check.Interval = "3s"
	check.HTTP = "http://172.18.32.111:8090/health"

	reg.Check = &check

	_ = ConsulClient.Agent().ServiceRegister(&reg)
}

func UnRegService() {
	ConsulClient.Agent().ServiceDeregister("ConsulService")
}
