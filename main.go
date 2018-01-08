package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/gin-contrib/cors.v1"
	"github.com/itnopadol/send_email/ctrl"
)

func main() {
	r := gin.New()
	r.Use(cors.Default())

	r.GET("/email", ctrl.SendEmail)
	r.Run(":8099")
}

