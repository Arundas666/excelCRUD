package main

import (
	"fmt"
	"vrwizards/pkg/config"
	"vrwizards/pkg/db"
	"vrwizards/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	fmt.Println("AA")
	db.SetupDatabase()
	db.SetupRedis()
	r := gin.Default()
	routes.Router(&r.RouterGroup)
	r.Run(":8080")

}
