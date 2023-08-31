package main

import (
	"fmt"
	"log"
	"net/http"

	"app/db"
)

var (
	ServiceName = "test-go-hong"
	port        = "3000"
)

func main() {
	db := db.ConnectDatabase()
	defer db.Close()
	fmt.Println("Table created successfully!", db)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(fmt.Sprintf("ok %s", ServiceName)))
	})
	fmt.Println("start service with port: ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}
