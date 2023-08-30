package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func init() {

}

var (
	ServiceName = "test-go-hong"
	port        = "8100"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(fmt.Sprintf("ok %s", ServiceName)))
	})
	fmt.Println("start service with port: ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

	db, err := sql.Open("mysql", "root:123@tcp(mysql:3306)/TestDB")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("Table created successfully!", db)

}
