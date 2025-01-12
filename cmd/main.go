package main

import "github.com/dyhalmeida/go-apis/configs"

func main() {
	config := configs.NewConfig()
	println(config.GetDatabaseDriver())
}
