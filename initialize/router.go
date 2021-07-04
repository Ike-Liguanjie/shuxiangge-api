package initialize

import (
	"github.com/gin-gonic/gin"
	"shuxiangge-api/router"
)

func Routers() * gin.Engine {
	var Router = gin.Default()
	PrivateGroup := Router.Group("/v1")
	{
		router.InitCategoryRouter(PrivateGroup)
	}
	return Router
}
