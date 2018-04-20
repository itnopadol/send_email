package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/itnopadol/send_email/ctrl"
	"gopkg.in/gin-contrib/cors.v1"
)

var (
	templates *template.Template
)

func main() {
	r := gin.New()
	r.Use(cors.Default())

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/email", ctrl.PaybillEmail)
	r.GET("/email/html", ctrl.ShowPaybillDocNo)

	r.Run(":20000")
}
