package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shuxiangge-api/global"
	"shuxiangge-api/model"
)

func GetCategoryList(c *gin.Context)  {
	var category []model.BooksCategory
	global.DB.Select("id", "name", "order").Find(&category).Scan(&category)
	c.JSON(http.StatusOK, category)
}
