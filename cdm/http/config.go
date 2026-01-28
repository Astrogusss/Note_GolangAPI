package main

import (
	"io"
	"log/slog"
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

func NewLogger(out io.Writer, minLevel slog.Level) *slog.Logger{
	return slog.New(slog.NewJSONHandler(out, &slog.HandlerOptions{
		AddSource: true,
		Level: minLevel,
	}))
}