package main

import (
	"fmt"
	"github.com/unclepepper/clean-go/configs"
	"github.com/unclepepper/clean-go/pkg/store/postgres"
)

func main() {
	conf := configs.LoadConfig()
	newDb := postgres.NewPostgres(conf)

	fmt.Println(newDb)
}
