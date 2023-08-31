package db

import (
	"app/models"
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func ConnectDatabase() (db *pg.DB) {
	db = pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "123",
		Database: "postgres",
	})
	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	f, err := os.Open("user.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Insert list users from csv
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	users := createUserList(data)
	for _, user := range users {
		_, err = db.Model(user).Insert()
		if err != nil {
			panic(err)
		}
	}

	return db
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*models.User)(nil),
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})

		if err != nil {
			return err
		}
	}
	return nil
}

func createUserList(data [][]string) []*models.User {
	var userList []*models.User

	for i, line := range data {
		if i > 0 {
			user := &models.User{}

			for j, field := range line {
				switch j {
				case 0:
					user.FullName = field
				case 1:
					emp_code, _ := strconv.ParseInt(field, 0, 64)
					user.EmployeeCode = int(emp_code)
				case 2:
					user.Email = field
				case 3:
					user.Role = field
				}
			}

			userList = append(userList, user)
		}
	}
	return userList
}
