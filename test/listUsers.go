package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type User struct {
	full_name     string
	employee_code int
	email         string
	role          string
}

func createFile(data [][]string) {

	fileName := "role" + ".csv"

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	writer.WriteAll(data)
}

func createUserList(data [][]string) []User {
	var userList []User
	for i, line := range data {
		if i > 0 {
			var user User
			for j, field := range line {
				switch j {
				case 0:
					user.full_name = field
				case 1:
					emp_code, err := strconv.ParseInt(field, 0, 64)
					if err != nil {
						fmt.Printf("Thiếu mã nhân viên thứ %d\n", i)
					}
					user.employee_code = int(emp_code)
				case 2:
					user.email = field
				case 3:
					user.role = field
				}
			}
			userList = append(userList, user)
		}
	}
	return userList
}

func main() {

	f, err := os.Open("user.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	userList := createUserList(data)

	fmt.Printf("%+v\n", userList[2])

}
