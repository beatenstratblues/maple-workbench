package main

import (
	"fmt"

	"github.com/beatenstratblues/maple-workbench/internals/config"
	"github.com/beatenstratblues/maple-workbench/internals/store"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		fmt.Println("There is a configuration error", err)
	}
	db, err, status := store.Connection("postgres://admin:123456789@127.0.0.1:5432/maple_dev?sslmode=disable", 25, 25, 5)
	if err != nil {
		fmt.Println(err)
	}
	db.Query("SELECT version()")
	fmt.Println(status)
	fmt.Println(conf)
	fmt.Println("Welcome to the mongoDB workbench!")
}
