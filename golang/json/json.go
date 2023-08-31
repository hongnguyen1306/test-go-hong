/*
  - Unmarshaling (Parsing từ JSON sang struct):
    +
  - Marshaling (Convert từ GO Object sang JSON)
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	// giải nén chuỗi JSON thành một struct sử dụng Unmarshal
	jsonString := `{
		"title":"Learning JSON",
		"author":"Lanka"
	}`
	var book Book
	err := json.Unmarshal([]byte(jsonString), &book)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", book)

	// Unstructured Data dùng map[string]interface
	jsonString2 := `{"title":"Learning JSON in Golang","author":"Lanka"}`
	var book2 map[string]interface{}
	err2 := json.Unmarshal([]byte(jsonString2), &book2)
	if err2 != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", book)

	//Marshalling JSON, dùng json.Marsal
	booking := Book{Title: "LEARNING GOLANG", Author: "ALAKA"}
	byteArr, err := json.Marshal(booking)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byteArr))
}
