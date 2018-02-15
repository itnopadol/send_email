package ctrl

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itnopadol/send_email/model"
)

func SendEmail(c *gin.Context) {
	c.Keys = headerKeys
	access_token := c.Request.URL.Query().Get("access_token")
	ar_code := c.Request.URL.Query().Get("ar_code")
	doc_no := c.Request.URL.Query().Get("doc_no")
	email := c.Request.URL.Query().Get("email")

	paybill := new(model.Paybill)
	err := paybill.PaybillEmail(access_token, ar_code, doc_no, email)
	fmt.Println("Ctrl Send Email ")

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = nil
		c.JSON(http.StatusOK, rs)
	}

}

func ShowPaybillDocNo(c *gin.Context) {
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	ar_code := c.Request.URL.Query().Get("ar_code")
	doc_no := c.Request.URL.Query().Get("doc_no")

	paybill := new(model.Paybill)
	p, err := paybill.ShowPaybillDocNo(dbc, ar_code, doc_no, access_token)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.HTML(http.StatusOK, "invoice.html", p)
}
