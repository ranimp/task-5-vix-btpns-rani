package main

import (
	"github.com/ranimp/task-5-vix-btpns-rani/database"
	"github.com/ranimp/task-5-vix-btpns-rani/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
