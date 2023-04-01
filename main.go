package main

import (
	"echo_golang/route"
	user_migrate "echo_golang/utils"
)

func main() {
	user_migrate.UserMigrate()

	e := route.Routers()
	e.Logger.Fatal(e.Start(":8000"))
}
