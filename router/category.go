package router

import (
	"github.com/gin-gonic/gin"
	"shuxiangge-api/api/v1"
)

func InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouter := Router.Group("category")
	CategoryRouter.GET("categoryList", v1.GetCategoryList)
}
