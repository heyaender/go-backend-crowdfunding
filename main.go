package main

import (
	"go-backend-crowdfunding/databases"
	"go-backend-crowdfunding/routes"
)

// It connects to the MySQL database and launches the application.
func main() {
	databases.MySQLConnect()
	routes.LaunchApp()
}
