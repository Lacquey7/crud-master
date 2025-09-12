package main

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	Port = "PORT_SERVER"
	Addr = "ADDRESS_SERVER"

	MqPort = "PORT_MQ"
	MqAddr = "ADDRESS_MQ"
	MqUser = "USERNAME_MQ"
	MqPass = "PASSWORD_MQ"

	ADDR_BILLING   = "BILLING_APP_SRV"
	PORT_BILLING   = "BILLING_APP_PORT"
	ADDR_INVENTORY = "INVENTORY_APP_SRV"
	PORT_INVENTORY = "INVENTORY_APP_PORT"
)

type Configuration struct {
	//Server
	Addr string
	Port string

	//RabbitMQ
	MqAddr string
	MqPort string
	MqUser string
	MqPass string

	//Services
	BillingAddr   string
	BillingPort   string
	InventoryAddr string
	InventoryPort string
}

func NewConfig() *Configuration {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	return &Configuration{
		//Server
		Addr: os.Getenv(Addr),
		Port: os.Getenv(Port),

		//RabbitMQ
		MqAddr: os.Getenv(MqAddr),
		MqPort: os.Getenv(MqPort),
		MqUser: os.Getenv(MqUser),
		MqPass: os.Getenv(MqPass),

		//Services
		BillingAddr:   os.Getenv(ADDR_BILLING),
		BillingPort:   os.Getenv(PORT_BILLING),
		InventoryAddr: os.Getenv(ADDR_INVENTORY),
		InventoryPort: os.Getenv(PORT_INVENTORY),
	}
}
