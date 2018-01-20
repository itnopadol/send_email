package main

import (
	"github.com/gin-gonic/gin"
//	"gopkg.in/gin-contrib/cors.v1"
//	"github.com/itnopadol/send_email/ctrl"
////	"net/http"
//	"net/http"
	"github.com/itnopadol/send_email/ctrl"
	//"net/http"
	"gopkg.in/gin-contrib/cors.v1"
)



func main() {
	r := gin.New()
	r.Use(cors.Default())

	r.GET("/email", ctrl.SendEmail)
	r.LoadHTMLGlob("templates/*")
	r.GET("/email/html",ctrl.X)
	r.Run(":8099")

}




