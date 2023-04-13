package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rikurunico/gin-basic/database"
	"github.com/rikurunico/gin-basic/route"
)

func main() {
	database.InitDB()

	router := gin.Default()

	route.SetupRoute(router)

	router.Run(":8080")
}
