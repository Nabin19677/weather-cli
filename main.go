/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"weather-cli/cmd"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	cmd.Execute()
}
