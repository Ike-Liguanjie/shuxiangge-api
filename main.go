package main

import (
	"shuxiangge-api/global"
	"shuxiangge-api/initialize"
)

func main()  {
	global.DB = initialize.Gorm()
	if global.DB != nil{
		db, _ :=global.DB.DB()
		defer db.Close()
	}
	Router := initialize.Routers()
	Router.Run()
}