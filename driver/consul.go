package driver

import (
	"github.com/hashicorp/consul/api"
)

var StockServiceAddr string

func GetService() {
	cfg := api.DefaultConfig()
	cfg.Address = "localhost:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	services, err := client.Agent().ServicesWithFilter(`Service == "stock_srv-1"`)
	if err != nil {
		panic(err)
	}
	for _, value := range services {
		StockServiceAddr = value.Address
	}
}
