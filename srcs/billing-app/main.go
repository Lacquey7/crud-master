package main

import (
	"billing-app/internal"
	"billing-app/internal/server"
	"billing-app/pkg"
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	// initialisation des configs
	cfg := NewConfig()
	println(fmt.Sprintf("Config loaded"))

	// initialisation DB
	db := pkg.NewDatabase(cfg.DBAddr, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
	conn := db.Connect()

	// initialisation de l'application
	app := internal.NewApp(conn)

	// initialisation RabbitMQ
	mq := pkg.NewRabbitMQ(cfg.MqAddr, cfg.MqPort, cfg.MqUser, cfg.MqPass, app.Handlers.Service)
	go mq.Connect()

	//initialisation des routes
	r := mux.NewRouter()
	app.Router.Router(r)

	// initialisation du serveur
	s := server.NewServer(cfg.Addr, cfg.Port, r)
	println(fmt.Sprintf("Server started at http://%s:%s", cfg.Addr, cfg.Port))
	s.Start()
}
