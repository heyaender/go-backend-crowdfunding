package main

import (
	"go-backend-crowdfunding/databases"
	"go-backend-crowdfunding/routes"
)

func main() {
	databases.MySQLConnect()
	routes.LaunchApp()
}
