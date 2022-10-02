package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(router *gin.Engine) {
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"up": true,
		})
	})
}
