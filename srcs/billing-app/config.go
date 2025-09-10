package main

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	Port = "PORT_SERVER"
	Addr = "ADDRESS_SERVER"

	DbUsername = "USERNAME_DB"
	DbPassword = "PASSWORD_DB"
	DbName     = "NAME_DB"
	DbAddr     = "ADDRESS_DB"
	DbPort     = "PORT_DB"

	MqPort = "PORT_MQ"
	MqAddr = "ADDRESS_MQ"
	MqUser = "USERNAME_MQ"
	MqPass = "PASSWORD_MQ"
)

type Configuration struct {
	//Server
	Addr string
	Port string

	//DB
	DBAddr string
	DBPort string
	DBUser string
	DBPass string
	DBName string

	//RabbitMQ
	MqAddr string
	MqPort string
	MqUser string
	MqPass string
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

		//Database
		DBAddr: os.Getenv(DbAddr),
		DBPort: os.Getenv(DbPort),
		DBUser: os.Getenv(DbUsername),
		DBPass: os.Getenv(DbPassword),
		DBName: os.Getenv(DbName),

		//RabbitMQ
		MqAddr: os.Getenv(MqAddr),
		MqPort: os.Getenv(MqPort),
		MqUser: os.Getenv(MqUser),
		MqPass: os.Getenv(MqPass),
	}
}
