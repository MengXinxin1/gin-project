package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"main.go/common"
	"main.go/controller"
	"main.go/middleware"
)

func main() {
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.POST("/register1/register", controller.Register)
	r.POST("/register1/login", controller.Login)
	panic(r.Run("localhost:8082"))
}
