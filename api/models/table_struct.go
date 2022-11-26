package models

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

type Attendee struct {
	User_id     int `json:"user_id"`
	Customer_id int `json:"customer_id"`
	A_id        int `gorm:"<-:update;primary_key" json:"a_id"`
}

type Countersign struct {
	C_id        int       `gorm:"<-:update;primary_key" json:"c_id"`
	Dep_id      int       `json:"dep_id"`
	Feedback    string    `json:"feedback"`
	Create_time time.Time `json:"create_time"`
	Status      bool      `json:"status"`
}

// type CountSign struct {
// 	C_id        int
// 	Dep_id      int
// 	Status      bool
// 	Feedback    string
// 	Create_time time.Time
// }

type Customer struct {
	C_id        int       `gorm:"<-:update;primary_key" json:"c_id"`
	Short_name  string    `json:"short_name"`
	Eng_name    string    `json:"eng_name"`
	Name        string    `json:"name"`
	Zip_code    string    `json:"zip_code"`
	Address     string    `json:"address"`
	Tel         string    `json:"tel"`
	Fax         string    `json:"fax"`
	Email       string    `json:"email"`
	Map         string    `json:"map"`
	Creater     string    `json:"creater"`
	Create_time time.Time `json:"create_time"`
}

type Customer_demand struct {
	C_id              int           `gorm:"<-:update;primary_key" json:"c_id"`
	Customer_name     string        `json:"customer_name"`
	Subject           string        `json:"subject"`
	Budget            string        `json:"budget"`
	Remarks           string        `json:"remarks"`
	Extend_type_id    int           `json:"extend_type_id"`
	Extend_rem        string        `json:"extend_rem"`
	Est_quantity      int           `json:"est_quantity"`
	Countersign_id    pq.Int64Array `json:"countersign_id" gorm:"type:integer[]"`
	Meeting_id        pq.Int64Array `json:"meeting_id" gorm:"type:integer[]"`
	Date_for_recive   time.Time     `json:"date_for_recive"`
	Task_id           int           `json:"task_id" `
	Accept            bool          `json:"accept"`
	Date_for_devlop   time.Time     `json:"date_for_devlop"`
	Eva_report        bool          `json:"eva_report"`
	Date_for_expected time.Time     `json:"date_for_expected"`
	Date_for_done     time.Time     `json:"date_for_done"`
	Project_code      string        `json:"project_code"`
	Salesman_name     string        `json:"salesman_name"`
	File_id           pq.Int64Array `json:"file_id" gorm:"type:integer[]"`
	Creater           int           `json:"creater"`
	Create_time       time.Time     `json:"create_time"`
	Code              string        `json:"code"`
	Status            string        `json:"status"`
}

type Department struct {
	D_id         int    `gorm:"<-:update;primary_key" json:"d_id"`
	Name         string `json:"name"`
	Eng_name     string `json:"eng_name"`
	Introduction string `json:"introduction"`
	Manager      int    `json:"manager"`
	Tel          string `json:"tel"`
	Fax          string `json:"fax"`
}

type File struct {
	F_id        int       `gorm:"<-:update;primary_key" json:"f_id"`
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	Create_time time.Time `json:"create_time"`
}

type Interview struct {
	I_id          int       `gorm:"<-:update;primary_key" json:"i_id"`
	Time          time.Time `json:"time"`
	Tpye          string    `json:"tpye"`
	Content       string    `json:"content"`
	Cus_demand_id int       `json:"cus_demand_id"`
	Create_time   time.Time `json:"create_time"`
	Creater       int       `json:"creater"`
}

type Jig_demand struct {
	J_id                      int           `gorm:"<-:update;primary_key" json:"j_id"`
	Kind                      string        `json:"kind"`
	Type                      string        `json:"type"`
	Urgent                    bool          `json:"urgent"`
	Title                     string        `json:"title"`
	Quantity                  int           `json:"quantity"`
	Date_for_notify           time.Time     `json:"date_for_notify"`
	Information               time.Time     `json:"information"`
	Expect_shipment_day       time.Time     `json:"expect_shipment_day"`
	Po_date                   time.Time     `json:"po_date"`
	Order_name                string        `json:"order_name"`
	Project_id                int           `json:"project_id"`
	Item                      string        `json:"item"`
	Standard                  string        `json:"standard"`
	Summary                   string        `json:"summary"`
	Remark                    string        `json:"remark"`
	Liaison                   string        `json:"liaison"`
	Liaison_phone             string        `json:"liaison_phone"`
	Date_for_demand           time.Time     `json:"date_for_demand"`
	Date_for_respond          time.Time     `json:"date_for_respond"`
	Date_for_respond_of_limit time.Time     `json:"date_for_respond_of_limit"`
	Customer_id               int           `json:"customer_id"`
	Create_time               time.Time     `json:"creater_time"`
	Creater                   int           `json:"creater"`
	Task_id                   pq.Int64Array `json:"task_id" gorm:"type:integer[]"`
	Countersign_id            pq.Int64Array `json:"countersign_id" gorm:";type:integer[]"`
	Meeting_id                pq.Int64Array `json:"meeting_id" gorm:";type:integer[]"`
}

type Logs struct {
	L_id        int       `gorm:"<-:update;primary_key" json:"l_id"`
	Type        string    `json:"type"`
	Tablename   string    `json:"tablename"`
	Sql_code    string    `json:"sql_code"`
	Content     string    `json:"content"`
	User        string    `json:"user"`
	Create_time time.Time `json:"creater_time"`
}

type Machine_type struct {
	M_id        int       `gorm:"<-:update;primary_key" json:"m_id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Create_time time.Time `json:"creater_time"`
}

type Machine_work_place struct {
	M_id            int       `gorm:"<-:update;primary_key" json:"m_id"`
	Enable          bool      `json:"enable"`
	Place           string    `json:"place"`
	Machine_type_id int       `json:"machine_type_id"`
	Pn_code         string    `json:"pn_code"`
	Sn_code         string    `json:"sn_code"`
	Customer_id     int       `json:"customer_id"`
	Project_code    string    `json:"project_code"`
	Remark          string    `json:"remark"`
	Create_time     time.Time `json:"create_time"`
	Creater         int       `json:"creater"`
}

type Manufacture_order struct {
	M_id                int           `gorm:"<-:update;primary_key" json:"m_id"`
	Customer            string        `json:"customer_id"`
	Order_name          string        `json:"order_name"`
	Amount              string        `json:"amount"`
	Shipment_location   string        `json:"shipment_location"`
	Open_date           time.Time     `json:"open_date"`
	Close_date          time.Time     `json:"close_date"`
	Expect_shipment_day time.Time     `json:"expect_shipment_day"`
	Sales_assistant     string        `json:"sales_assistant"`
	Recipient           string        `json:"recipient"`
	Contact_person      string        `json:"contact_person"`
	Remarks             string        `json:"remarks"`
	Create_time         time.Time     `json:"create_time"`
	Project_id          int           `json:"project_id"`
	Copy                pq.Int64Array `json:"copy" gorm:"type:integer[]"`
	Status              string        `json:"status"`
	Creater             int           `json:"creater"`
}

type Meeting struct {
	M_id        int       `gorm:"<-:update;primary_key" json:"m_id"`
	Name        string    `json:"name"`
	Room        string    `json:"room"`
	Start_date  time.Time `json:"start_date"`
	End_date    time.Time `json:"end_date"`
	Create_time time.Time `json:"creater_time"`
	Chairman    int       `json:"chairman"`
	Attendees   int       `json:"attendees"`
}

type Notice_time struct {
	N_id      int       `gorm:"<-:update;primary_key" json:"n_id"`
	Meet_id   int       `json:"meet_id"`
	Meet_time time.Time `json:"meet_time"`
}

type Project struct {
	P_id                   int            `gorm:"<-:update;primary_key" json:"p_id"`
	Code                   string         `json:"code"`
	Name                   string         `json:"name"`
	Customer_name          string         `json:"customer_name"`
	Date_for_start         string         `json:"date_for_start"`
	Date_for_end           string         `json:"date_for_end"`
	Salesman_name          string         `json:"salesman_name"`
	Serviceman_name        string         `json:"serviceman_name"`
	Projectman_name        string         `json:"projectman_name"`
	Status                 string         `json:"status"`
	Create_time            time.Time      `json:"create_time"`
	Type                   string         `json:"type"`
	Project_member         pq.StringArray `json:"project_member" gorm:"type:text[]"`
	Meeting_id             pq.StringArray `json:"meeting_id" gorm:"type:text[]"`
	File                   pq.StringArray `json:"file"  gorm:"type:text[]"`
	Task_id                int            `json:"task_id" `
	Manufactrue_order_list pq.StringArray `json:"manufactrue_order_list"  gorm:"type:text[]"`
	Date_for_pay           time.Time      `json:"date_for_pay"`
	Date_for_delivery      time.Time      `json:"date_for_delivery"`
	Date_for_check         time.Time      `json:"date_for_check"`
	Inner_id               int            `json:"inner_id"`
}

type Sysuser struct {
	Logonid    string `gorm:"<-:update;primary_key" json:"logonid"`
	Name       string `json:"name"`
	Name1      string `json:"name1"`
	Title      string `json:"title"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	Empno      string `json:"empno"`
	Logenabled string `json:"logenabled"`
	Labuser    string `json:"labuser"`
	Gboss      string `json:"gboss"`
	Sysboss    string `json:"sysboss"`
	Gkind      string `json:"gkind"`
	Gkindboss  string `json:"gkindboss"`
	Rdopt      string `json:"rdopt"`
}

type Task struct {
	T_id            int       `gorm:"<-:update;primary_key" json:"t_id"`
	Type            string    `json:"type"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Principal       string    `json:"principal"`
	Before_id       string    `json:"before_id"`
	Time_for_start  time.Time `json:"time_for_start"`
	Time_for_done   time.Time `json:"time_for_done"`
	Include_weekend bool      `json:"include_weekend"`
	Labor_hour      int       `json:"labor_hour"`
	Finish          bool      `json:"finish"`
	Creater         int       `json:"creater"`
	Create_time     time.Time `json:"create_time"`
}

type Work_item struct {
	W_id           int       `gorm:"<-:update;primary_key" json:"w_id"`
	Name           string    `json:"name"`
	Type           string    `json:"type"`
	Required       bool      `json:"required"`
	Provide_amount bool      `json:"provide_amount"`
	Create_time    time.Time `json:"create_time"`
}

type Worker_set struct {
	W_id              int       `gorm:"<-:update;primary_key" json:"w_id"`
	Work_item         int       `json:"work_item"`
	User_id           string    `json:"user_id"`
	Date_for_expected time.Time `json:"date_for_expected"`
	Date_for_done     time.Time `json:"date_for_done"`
	Amount            string    `json:"amount"`
	Remarks           string    `json:"remarks"`
	Create_time       time.Time `json:"create_time"`
}

type Auth struct {
	A_id        int        `gorm:"<-:update;primary_key" json:"a_id"`
	Name        string     `json:"name"`
	Path        string     `json:"path"`
	File        string     `json:"file"`
	Create_time *time.Time `json:"create_time"`
}

type Test_task struct {
	T_id   int            `gorm:"<-:update;primary_key" json:"t_id"`
	Detail postgres.Jsonb `gorm:"default:'{}'" json:"detail"`
}

type Labor_hour struct {
	H_id           int        `gorm:"<-:update;primary_key" json:"h_id"`
	Type           string     `json:"type"`
	Name           string     `json:"name"`
	Content        string     `json:"content"`
	Nature         string     `json:"nature"`
	Date_for_labor *time.Time `json:"date_for_labor"`
	Time_start     int        `json:"time_start"`
	Time_end       int        `json:"time_end"`
	Hour           int        `json:"hour"`
	Create_time    *time.Time `json:"create_time"`
}
