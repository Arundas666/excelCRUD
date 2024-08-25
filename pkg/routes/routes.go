package routes

import (
	"vrwizards/pkg/handlers"

	"github.com/gin-gonic/gin"
)
func Router(r *gin.RouterGroup){

	r.POST("/import", handlers.ImportData)
	r.DELETE("/delete",handlers.DeleteData)
    r.GET("/data", handlers.GetData)
    r.PUT("/data/:id", handlers.UpdateData)
}