package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"bytes"
	"net/smtp"
	"log"
	"html/template"
	//"net/http"
	//"path"
	//"image"
	//"image/jpeg"
	//"strconv"
	"net/http"
)

type Paybill struct{
	ArCode string `json:"ar_code" db:"ArCode"`
	ArName string `json:"ar_name" db:"ArName"`
	BillAddress string `json:"bill_address" db:"BillAddress"`
	BillSub []*Bill `json:"bill_sub"`
}

type Bill struct{
	DocNo string `json:"doc_no" db:"DocNo"`
	DocDate string `json:"doc_date" db:"DocDate"`
}

type Details struct{
	InvoiceNo string `json:"invoice_no" db:"InvoiceNo"`
	InvoiceDate string `json:"invoice_date" db:"InvoiceDate"`
	InvBalance float64 `json:"inv_balance" db:"InvBalance"`
}

type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

func (p *Paybill)TestEmail(db *sqlx.DB) (paybills []*Paybill , err error){
	sql := `select top 1 ArCode,b.name1 as ArName,b.BillAddress from dbo.bcpaybill a inner join dbo.bcar b on a.arcode = b.code where arcode = '0810244197'`
	fmt.Println("query = ", sql)
	err = db.Select(&paybills,sql)
	fmt.Println("query = ", sql, p.ArName)
	if err != nil {
		return nil, err
	}

	for _, pp := range paybills{
		sqlsub := `select DocNo,DocDate from dbo.bcpaybillsub where arcode = ?`
		fmt.Println("query sub= ", sqlsub, pp.ArCode)
		err = db.Select(&pp.BillSub, sqlsub, pp.ArCode)
		if err != nil {
			return nil, err
		}

		fmt.Println(pp.BillSub[0].DocNo)
	}



	subject := "Hello Email"
	receiver:= "it@nopadol.com"
	r := NewRequest([]string{receiver}, subject)
	r.Send("templates/paybill.html",paybills)//map[string]string{"username": "satit","surname" : "chomwattana"})

	return nil, nil
}

func NewRequest(to []string, subject string) *Request {
	return &Request{
		to:      to,
		subject: subject,
	}
}

func (r *Request) Send(templateName string, items interface{}) {
	err := r.parseTemplate(templateName, items)
	if err != nil {
		log.Fatal(err)
	}
	if ok := r.sendEmail(); ok {
		log.Printf("Email has been sent to %s\n", r.to)
	} else {
		log.Printf("Failed to send the email to %s\n", r.to)
	}
}

var (
	templates *template.Template
)

func index(w http.ResponseWriter, r *http.Request) {

	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (r *Request) parseTemplate(fileName string, data interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return err
	}
	r.body = buffer.String()
	return nil
}

func (r *Request) sendEmail() bool {
	body := "To: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + MIME + "\r\n" + r.body
	SMTP := fmt.Sprintf("%s:%d", "smtp.gmail.com", 587)
	if err := smtp.SendMail(SMTP, smtp.PlainAuth("", "nopadol_mailauto@nopadol.com", "[vdw,jwfh2012", "smtp.gmail.com"), "satit@nopadol.com", r.to, []byte(body)); err != nil {
		return false
	}
	return true
}

//func writeImage(w http.ResponseWriter, img *image.Image) {
//
//	buffer := new(bytes.Buffer)
//	if err := jpeg.Encode(buffer, *img, nil); err != nil {
//		log.Println("unable to encode image.")
//	}
//
//	w.Header().Set("Content-Type", "image/jpeg")
//	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
//	if _, err := w.Write(buffer.Bytes()); err != nil {
//		log.Println("unable to write image.")
//	}
//}