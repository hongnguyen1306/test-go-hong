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

func createUserList(data [][]string, w1, w2, w3 *csv.Writer) []User {
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

			switch user.role {
			case "MEMBER":
				if err := w2.Write(line); err != nil {
					log.Fatalln("error writing record to member file", err)
				}
			case "SUPER_ADMIN":
				if err := w1.Write(line); err != nil {
					log.Fatalln("error writing record to admin file", err)
				}
			case "ACCOUNTING":
				if err := w3.Write(line); err != nil {
					log.Fatalln("error writing record to accounting file", err)
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

	f1, _ := os.Create("admin.csv")
	defer f1.Close()
	f2, _ := os.Create("member.csv")
	defer f2.Close()
	f3, _ := os.Create("accounting.csv")
	defer f3.Close()

	w1 := csv.NewWriter(f1)
	defer w1.Flush()
	w2 := csv.NewWriter(f2)
	defer w2.Flush()
	w3 := csv.NewWriter(f3)
	defer w3.Flush()

	userList := createUserList(data, w1, w2, w3)

	fmt.Println(userList)
}
