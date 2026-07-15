package main

import (
	"fmt"
	"github.com/beatenstratblues/maple-workbench/internals/config"
)

func main() {
	conf, err := config.Load()
	if err!=nil {
		fmt.Println("There is a configuration error", err)
	}
	fmt.Println(conf)
	fmt.Println("Welcome to the mongoDB workbench!")
}
