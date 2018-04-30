package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//Config structure
type Config struct {
	CRMServer      string `json:"crmServer"`
	CRMUser        string `json:"crmUser"`
	CRMPwd         string `json:"crmPwd"`
	CRMDb          string `json:"crmDb"`
	SageServer     string `json:"sageServer"`
	SageUser       string `json:"sageUser"`
	SagePwd        string `json:"sagePwd"`
	SageDb         string `json:"sageDb"`
	SendgridAPIkey string `json:"sendgridAPIkey"`
}

type quote struct {
	QuoteID        int
	CustomerName   string
	ContactName    string
	BusinessID     int
	DepartmentName string
	QuoteValue     string
	JobNo          string
	Description    string
	DateDespatched string
}

type data struct {
	Name     string
	AreaName string
	Quotes   []quote
	Date     string
}

//Salesrep struct
type Salesrep struct {
	UserID int    `json:"userID"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Areas  []int  `json:"areas"`
}

var config Config
var dbCRM *sql.DB
var dbSage *sql.DB
var logfile *os.File

func init() {

	var err error

	//Create log.txt
	logfile, err = os.Create("proglog.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}

	fmt.Fprintln(logfile, time.Now(), " - log file successfully created")

	//Load application configuration from settings file
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Fprintln(logfile, time.Now(), " - config json file loaded")

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(logfile, time.Now(), " - config file decoded")

	//Connect to CRM database and test connection
	crmConnection := fmt.Sprintf("Server=%s;User ID=%s;Password=%s;database=%s;",
		config.CRMServer,
		config.CRMUser,
		config.CRMPwd,
		config.CRMDb)

	dbCRM, err = sql.Open("mssql", crmConnection)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(logfile, time.Now(), " - connected to CRM database")

	if err = dbCRM.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(logfile, time.Now(), " - CRM db ping")

	//Connect to Sage database and test conntection
	sageConnection := fmt.Sprintf("Server=%s;User ID=%s;Password=%s;database=%s;",
		config.SageServer,
		config.SageUser,
		config.SagePwd,
		config.SageDb)

	dbSage, err = sql.Open("mssql", sageConnection)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(logfile, time.Now(), " - connected to sage database")

	if err = dbSage.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(logfile, time.Now(), " - sageDB ping")

}

func main() {

	fmt.Fprintln(logfile, time.Now(), " - func main initialised")

	var areaID int
	var areaName string
	var sageDb string
	var areaOutput string

	jsonFile, err := os.Open("salesreps.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	fmt.Println("Successfully opened salesreps.json")

	fmt.Fprintln(logfile, time.Now(), " - opened salesreps.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var salesReps []Salesrep

	json.Unmarshal(byteValue, &salesReps)

	salesAreas, err := dbCRM.Query("SELECT area_id, area_name FROM tbl_area")
	if err != nil {
		log.Fatal(err)
	}
	defer salesAreas.Close()

	fmt.Fprintln(logfile, time.Now(), " - sales areas obtained from CRM db")

	areaMap := make(map[int]string)
	for salesAreas.Next() {
		err := salesAreas.Scan(&areaID, &areaName)
		if err != nil {
			log.Fatal(err)
		}
		areaMap[areaID] = strings.TrimSpace(areaName)
	}

	fmt.Fprintln(logfile, time.Now(), " - area map created")

	for _, salesRep := range salesReps {

		fmt.Println(salesRep.Name)

		fmt.Fprintln(logfile, time.Now(), "  salesrep names obtained from json file")

		areaList := make([]string, 0)
		areaFilter := ""
		for _, areaID := range salesRep.Areas {
			areaList = append(areaList, areaMap[areaID])
			areaFilter += strconv.Itoa(areaID) + ","
		}
		if strings.HasSuffix(areaFilter, ",") {
			areaFilter = areaFilter[:len(areaFilter)-1]
		}

		fmt.Fprintln(logfile, time.Now(), " - arealist created")

		sql1 := `SELECT c.customer_name, t.contact_name, q.quote_id, a.business_id, d.department_name, q.quote_value, COALESCE(q.job_no, '') AS job_no
				 FROM tbl_customer c
					JOIN tbl_site s ON c.customer_id = s.customer_id
					JOIN tbl_site_area sa on s.site_id = sa.site_id
					JOIN tbl_area a ON sa.area_id = a.area_id
					JOIN tbl_contact t ON sa.site_id = t.site_id AND sa.area_id = t.area_id
					JOIN tbl_quote q ON t.contact_id = q.contact_id
					JOIN tbl_department d ON q.department_id = d.department_id
				 WHERE q.quote_date = CONVERT(NVARCHAR(11), CURRENT_TIMESTAMP, 106)
		 		 AND sa.area_id IN (` + areaFilter + `)
				 ORDER BY q.quote_id`

		quoteRows, err := dbCRM.Query(sql1)
		if err != nil {
			log.Fatal(err)
		}
		defer quoteRows.Close()

		fmt.Fprintln(logfile, time.Now(), " - sql1 executed")

		quotes := make([]quote, 0)

		for quoteRows.Next() {

			var q quote
			var valueFloat float64

			err := quoteRows.Scan(&q.CustomerName, &q.ContactName, &q.QuoteID, &q.BusinessID, &q.DepartmentName, &valueFloat, &q.JobNo)
			q.QuoteValue = fmt.Sprintf("%.2f", valueFloat)

			if err != nil {
				log.Fatal(err)
			}

			if q.JobNo != "" {

				if q.BusinessID == 1 {
					sageDb = "rjw"
				} else {
					sageDb = "chr"
				}

				sql1 = `SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
						SELECT j.description1, COALESCE(CONVERT(NVARCHAR(11), soh.date_despatched, 106), '-')
						FROM ` + sageDb + `.scheme.jcmastm j
						  JOIN ` + sageDb + `.scheme.opheadm soh ON j.job_code = soh.order_no
						WHERE j.[job_code] = ?`

				err = dbSage.QueryRow(sql1, q.JobNo).Scan(&q.Description, &q.DateDespatched)
				if err != nil && err != sql.ErrNoRows {
					log.Fatal(err)
				}
				fmt.Fprintln(logfile, time.Now(), " - sql2 executed")
				q.Description = strings.TrimSpace(q.Description)

			} else {
				q.JobNo = "Not supplied"
				q.Description = "-"
				q.DateDespatched = "-"
			}

			quotes = append(quotes, q)

			fmt.Fprintln(logfile, time.Now(), " - quotes append worked")

		}

		if len(areaList) == 1 {
			areaOutput = "area " + areaList[0] + " only"
		} else if len(areaList) == 2 {
			areaOutput = "areas " + areaList[0] + " and " + areaList[1]
		} else if len(areaList) >= 3 {
			areaOutput = "areas "
			for i := 0; i < len(areaList)-1; i++ {
				areaOutput += areaList[i] + ", "
			}
			if strings.HasSuffix(areaOutput, ", ") {
				areaOutput = areaOutput[:len(areaOutput)-2]
			}
			areaOutput += " and " + areaList[len(areaList)-1]
		}

		todaysDate := time.Now().Format("2 Jan 2006")

		t, _ := template.ParseFiles("Quotationsummary.tpl")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(logfile, time.Now(), " - quotationsummary.tpl parsed")

		//Send email via sendgrid
		var email bytes.Buffer

		err = t.Execute(&email, data{salesRep.Name, areaOutput, quotes, todaysDate})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(logfile, time.Now(), " - t.Execute successful")

		m := mail.NewV3Mail()
		m.SetFrom(mail.NewEmail("Rewinds & J Windsor Ltd", "donotreply@rjweng.com"))
		m.Subject = fmt.Sprintf("Quotation Report (%s) - %s", salesRep.Name, todaysDate)

		p := mail.NewPersonalization()
		tos := []*mail.Email{
			mail.NewEmail(salesRep.Name, salesRep.Email),
		}
		p.AddTos(tos...)

		ccs := []*mail.Email{
			mail.NewEmail("Lee Windsor", "lee@rjweng.com"),
			mail.NewEmail("John Windsor", "john.windsor@rjweng.com"),
			mail.NewEmail("Luke Windsor", "luke.windsor@rjweng.com"),
			mail.NewEmail("Fraser Whittle", "fraser@rjweng.com"),
			mail.NewEmail("Leigh Jenkins", "leigh@marinconsulting.co.uk"),
		}
		p.AddCCs(ccs...)

		m.AddPersonalizations(p)
		m.AddContent(mail.NewContent("text/html", email.String()))

		request := sendgrid.GetRequest(config.SendgridAPIkey, "/v3/mail/send", "https://api.sendgrid.com")
		request.Method = "POST"
		request.Body = mail.GetRequestBody(m)

		response, err := sendgrid.API(request)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Body)
			fmt.Println(response.Headers)
		}

	}

	fmt.Fprintln(logfile, time.Now(), " - emails sent")

}
