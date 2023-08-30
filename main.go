package main

import (
	"fmt"
	"log"
	"net/http"
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
}
