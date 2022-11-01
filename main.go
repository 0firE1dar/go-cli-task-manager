package main

import (
	"taskmanager/cmd"
	"taskmanager/database"
)

func main() {

	database.Connect()
	defer database.Close()

	cmd.Execute()
}