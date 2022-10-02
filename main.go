package main

import (
	"github.com/gin-gonic/gin"
	"redfox/db"
	"redfox/route"
	"redfox/telemetry"
)

func main() {
	telemetry.Init()
	db.Connect()

	router := gin.Default()
	router.SetTrustedProxies(nil)

	route.Index(router)
	route.User(router)

	router.Run()
}
