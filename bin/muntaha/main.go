package main

import (
	"log"
	"os"

	"github.com/minhajuddinkhan/muntaha"
	"github.com/minhajuddinkhan/muntaha/cli"
)

var (
	dbPassword = os.Getenv("DB_PWD")
	dbUser     = os.Getenv("DB_USER")
	dbPort     = os.Getenv("DB_PORT")
	dbHost     = os.Getenv("DB_HOST")
	dbName     = os.Getenv("DB_NAME")
	appHost    = os.Getenv("HOST")
	appPort    = os.Getenv("PORT")
)

func main() {

	conf := muntaha.Configuration{
		DB: muntaha.DBConf{
			Host:     dbHost,
			Port:     dbPort,
			User:     dbUser,
			Password: dbPassword,
		},
		HttpPort: appPort,
		HttpHost: appHost,
	}

	err := cli.Run(conf)
	if err != nil {
		log.Fatal(err)
	}

}
