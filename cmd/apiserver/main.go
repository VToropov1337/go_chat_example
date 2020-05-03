package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"go_rest/internal/app/apiserver"
	"log"
)

var (
	configPath string
	)


func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}


func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	fmt.Println(config)
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}



//migrate create -ext sql -dir migrations create_users
//migrate -path migrations -database "postgres://localhost/restapi_dev?sslmode=disable" up
