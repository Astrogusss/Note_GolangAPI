package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	Server_Port string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil{panic(err)}

	config := Config{}
	
	//como so temos somente uma configuracao no env, vamos colocar somente um campo na struct
	if aux, err := os.LookupEnv("SERVER_PORT"); !err {
		panic(err)
	} else {
		config.Server_Port = aux
	}

	return &config
}