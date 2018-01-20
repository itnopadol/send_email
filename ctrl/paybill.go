package ctrl

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/itnopadol/send_email/model"
)

func SendEmail(c *gin.Context){
	c.Keys = headerKeys

	paybill := new(model.Paybill)
	//err := c.BindJSON(paybill)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	p, err := paybill.TestEmail(dbc)
	fmt.Println("Ctrl Send Email ")

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : "+err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = p
		c.JSON(http.StatusOK, rs)
	}

}


//func ShowEmail(c *gin.Context){
//	c.Keys = headerKeys
//
//	paybill := new(model.Paybill)
//	err := paybill.ShowEmail(dbc)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	fmt.Println("Ctrl Send Email ")
//}

func X(c *gin.Context) {
		c.Keys = headerKeys

		paybill := new(model.Paybill)
		p, err := paybill.ShowEmail(dbc)
		if err != nil {
			fmt.Println(err.Error())
		}
		c.HTML(http.StatusOK, "test.html", p)

}

