/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"weather-cli/cmd"
	"weather-cli/utils"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	cmd.Execute()
}

func init() {
	utils.CreateFolderT()
}
