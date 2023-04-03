package main

import (
	"echo_golang/route"
	"echo_golang/utils"
)

func main() {
	utils.UserMigrate()
	e := route.Routers()
	e.Logger.Fatal(e.Start(":8000"))
}
