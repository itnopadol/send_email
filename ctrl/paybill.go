package ctrl

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/itnopadol/send_email/model"
)

func SendEmail(c *gin.Context){
	c.Keys = headerKeys

	fmt.Println("Ctrl Send Email ")

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
