package model

import (
	// "bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	// "os"

	"github.com/jmoiron/sqlx"
)

type Paybill struct {
	ArCode        string         `json:"ar_code" db:"ArCode"`
	ArName        string         `json:"ar_name" db:"ArName"`
	TaxNo         string         `json:"tax_no" db:"TaxNo"`
	BillAddress   string         `json:"bill_address" db:"BillAddress"`
	ArDebtBalance float64        `json:"ar_debt_balance" db:"ArDebtBalance"`
	DebtLimit1    float64        `json:"debt_limit_1" db:"DebtLimit1"`
	DebtLimitBal  float64        `json:"debt_limit_bal" db:"DebtLimitBal"`
	DebtAmount    float64        `json:"debt_amount" db:"DebtAmount"`
	ChqOnHand     float64        `json:"chq_on_hand" db:"ChqOnHand"`
	ChqReturn     float64        `json:"chq_return" db:"ChqReturn"`
	DocNo         string         `json:"doc_no" db:"DocNo"`
	DocDate       string         `json:"doc_date" db:"DocDate"`
	DueDate       string         `json:"due_date" db:"DueDate"`
	SumOfInvoice  float64        `json:"sum_of_invoice" db:"SumOfInvoice"`
	Subs          []*PaybillSub  `json:"invoice_sub"`
	Balance       []*BillBalance `json:"balance"`
}

type PaybillSub struct {
	Id          int64   `json:"id" db:"Id"`
	InvoiceNo   string  `json:"invoice_no" db:"InvoiceNo"`
	InvoiceDate string  `json:"invoice_date" db:"InvoiceDate"`
	DueDateSub  string  `json:"due_date_sub" db:"DueDate"`
	InvBalance  float64 `json:"inv_balance" db:"InvBalance"`
	PayAmount   float64 `json:"pay_amount" db:"PayAmount"`
	PayBalance  float64 `json:"pay_balance" db:"PayBalance"`
	ItemName    string  `json:"item_name" db:"ItemName"`
	LineNumber  int     `json:"line_number" db:"LineNumber"`
}

type BillBalance struct {
	RowNumber    int64   `json:"row_number" db:"RowNumber"`
	MonthName    string  `json:"month_name" db:"MonthGroup"`
	MonthBalance float64 `json:"month_balance" db:"vSumBalance"`
	ArCode       string  `json:"ar_code" db:"ArCode"`
	ArName       string  `json:"ar_name" db:"ArName"`
	BillAddress  string  `json:"bill_address" db:"BillAddress"`
	Telephone    string  `json:"telephone" db:"Telephone"`
	Fax          string  `json:"fax" db:"Fax"`
}

type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; image/png; charset=\"UTF-8\";\n\n"
)

func (p *Paybill) PaybillEmail(access_token string, ar_code string, doc_no string, email string) error {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// text, _ := reader.ReadString('\n')
	// fmt.Println(text)

	// fmt.Println("Enter text: ")
	// text2 := ""
	// fmt.Scanln(text2)
	// fmt.Println(text2)

	// ln := ""
	// fmt.Sscanln("%v", ln)
	// fmt.Println(ln)

	subject := "Send PayBill"
	receiver := email
	r := NewRequest([]string{receiver}, subject)
	r.body = "http://localhost:8099/email/html?ar_code=" + ar_code + "&doc_no=" + doc_no
	body := "To: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + MIME + "\r\n" + r.body
	SMTP := fmt.Sprintf("%s:%d", "smtp.gmail.com", 587)
	if err := smtp.SendMail(SMTP, smtp.PlainAuth("", "nopadol_mailauto@nopadol.com", "[vdw,jwfh2012", "smtp.gmail.com"), "satit@nopadol.com", r.to, []byte(body)); err != nil {
		return err
	}
	return nil
}

func NewRequest(to []string, subject string) *Request {
	return &Request{
		to:      to,
		subject: subject,
	}
}

func (r *Request) Send(templateName string, items interface{}, ar_code string, doc_no string) {
	err := r.parseTemplate(templateName, items)
	if err != nil {
		log.Fatal(err)
	}
	if ok := r.sendEmail(ar_code, doc_no); ok {
		log.Printf("Email has been sent to %s\n", r.to)
	} else {
		log.Printf("Failed to send the email to %s\n", r.to)
	}
}

var (
	templates *template.Template
)

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

func (r *Request) sendEmail(ar_code string, doc_no string) bool {
	r.body = "http://localhost:8099/email/html?ar_code=" + ar_code + "&doc_no=" + doc_no
	body := "To: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + MIME + "\r\n" + r.body
	SMTP := fmt.Sprintf("%s:%d", "smtp.gmail.com", 587)
	if err := smtp.SendMail(SMTP, smtp.PlainAuth("", "nopadol_mailauto@nopadol.com", "[vdw,jwfh2012", "smtp.gmail.com"), "satit@nopadol.com", r.to, []byte(body)); err != nil {
		return false
	}
	return true
}

func (p *Paybill) ShowPaybillDocNo(db *sqlx.DB, ar_code string, doc_no string) (paybills []*Paybill, err error) {

	sql := `exec dbo.USP_API_ArDebtBalacnce ?, ?`
	fmt.Println("query sql = ", sql, ar_code, doc_no)
	err = db.Select(&paybills, sql, ar_code, doc_no)
	if err != nil {
		return nil, err
	}

	for _, pp := range paybills {
		sqlsub := `select InvoiceNo,rtrim(day(a.InvoiceDate))+'/'+rtrim(month(a.InvoiceDate))+'/'+rtrim(year(a.InvoiceDate)) as InvoiceDate,InvBalance,InvBalance,PayBalance,rtrim(day(b.DueDate))+'/'+rtrim(month(b.DueDate))+'/'+rtrim(year(b.DueDate)) as DueDate,LineNumber+1 as LineNumber,(select top 1 itemname from dbo.bcarinvoicesub where arcode = a.arcode and docno = a.invoiceno and docdate = a.invoicedate order by netamount desc) as ItemName from	dbo.bcpaybillsub a inner join dbo.bcpaybill c on a.docno = c.docno and a.arcode = c.arcode inner join dbo.bcarinvoice b on a.arcode = b.arcode and a.invoiceno = b.docno and a.InvoiceDate = b.docdate where	a.arcode = ? and a.docno = ? and c.billstatus = 0 and a.iscancel = 0 and c.iscancel = 0`
		fmt.Println("query sqlsub= ", sqlsub, pp.ArCode, pp.DocNo)
		err = db.Select(&pp.Subs, sqlsub, pp.ArCode, pp.DocNo)
		if err != nil {
			return nil, err
		}

		sqlbal := `exec dbo.USP_CD_ConfirmSaleOrderPayBill ?`
		fmt.Println("query sqlbal= ", sqlbal, pp.ArCode)
		err = db.Select(&pp.Balance, sqlbal, pp.ArCode)
		if err != nil {
			return nil, err
		}
	}

	return paybills, err
}
