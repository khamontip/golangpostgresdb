package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	//"github.com/rs/cors"

	_ "github.com/alexbrainman/odbc" // ODBC driver
	_ "github.com/lib/pq"
	//or
	//"github.com/jackc/pgx/v5"
)
var db *sql.DB
	var err error
type Product struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
	// Qty int `json: "qty"`
	// FlagDelete string `json:"flagdelete`
	// CreatedBy string `json:"createdby`
	// UpdatedBy string `json:"updatedeby"`
	// CreatedDate time.Time `json:createddate`
	// UpdatedDate time.Time `json:updateddate`
}
// สร้างโครงสร้างข้อมูลสำหรับส่งกลับไปให้ React
type ResponseData struct {
	Message string `json: "message"`
	Status string `json: "status"` 
}
const (
	host = "localhost"
	port = "5432"
	user = "postgres"
	password = "Passw0rd"
	dbname = "postgres"
)
func main() {
fmt.Println("----------Start--------")
connDB18()
fmt.Println("----------End--------")
}

func connDB18() {
	// 10 Oct Tuk Note can connect to postgres Database
	fmt.Println("Start connDB1-----")
psqlInfo := fmt.Sprintf("host=%s port=%s user=%s  "+
"password=%s dbname=%s sslmode=disable",host, port , user, password, dbname)
//-----------
//		dsn :="DSN=PostgreSQL;Uid=postgres;Pwd=Passw0rd"
// db, err = sql.Open("odbc", dsn)
// Open connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Can not open connection : %v",err)
	}
	defer db.Close()

	// test connection
	if err = db.Ping(); err != nil {
		log.Fatalf("DB Connected : %v", err)
	}
	http.HandleFunc("/api/products", createProductHandler)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
func createProductHandler(w http.ResponseWriter , r *http.Request) {
	fmt.Println("createProductHandler Start-----------")
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("ok"))
}
