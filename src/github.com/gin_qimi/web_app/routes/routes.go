package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin_qimi/web_app/logger"
	"github.com/gin_qimi/web_app/settings"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
	})
	return r
}
