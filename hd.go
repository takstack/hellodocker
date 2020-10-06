package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //comment
	"html"
	"log"
	"net/http"
	"strconv"
	//"strings"
	//"os"
	"time"
)

var db *sql.DB

func init() {

	var err error
	db, err = sql.Open("mysql", "ran:MysqlRand1@tcp(mysql:3306)/dockertest?parseTime=true")
	handleError(err)
	db.SetMaxOpenConns(11)
	db.SetMaxIdleConns(11)
	db.SetConnMaxLifetime(time.Second * 11 * 3)
	log.Println("Database Opened")
}

func seltest() []int {
	qry := "select num from sample;"
	res, err := db.Query(qry)
	handleError(err)
	var sl []int
	for res.Next() {
		var i int
		// for each row, scan the result into our tag composite object
		err := res.Scan(&i)
		handleError(err)

		sl = append(sl, i)
	}
	res.Close()
	return sl
}
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//log.SetOutput(os.Stdout)
		log.Println("in root")
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		log.Println("in hi")
		sl := seltest()
		fmt.Fprintf(w, "Hi\n\n")
		for _, elem := range sl {
			i := strconv.Itoa(elem)
			fmt.Fprintf(w, i+"\n\n")
		}

	})

	log.Fatal(http.ListenAndServe(":9800", nil))

}
func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
