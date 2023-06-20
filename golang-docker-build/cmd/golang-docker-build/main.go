package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	_ "modernc.org/sqlite"
)
// CSV file column name, ToDo: Automate the column name given in the CSV file
type emp_value struct {
	emp_id   string
	emp_name string
	dob      string
	role     string
	dept     string
}
// To get MYSQL connection
func getMySQLDB() *sql.DB {
	db, err := sql.Open("sqlite", "./sqlite_employees.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
// Function will import CSV data into SQLite DB
func main() {
	var db = getMySQLDB()
	employees := []emp_value{}
	file, err := os.Open("employees.csv")
	if err != nil {
		log.Fatal(err)
	}
	df := csv.NewReader(file)
	data, _ := df.ReadAll()
	for _, value := range data {
		employees = append(employees, emp_value{emp_id: value[0], emp_name: value[1], dob: value[2], role: value[3], dept: value[4]})
	}
	for i := 1; i < len(employees); i++ {
		emp_id, _ := strconv.Atoi(employees[i].emp_id)
		fmt.Println(db.Exec("insert into employee_details (emp_id, emp_name, emp_dob, emp_role, emp_dept) values(?,?,?,?,?)", emp_id, employees[i].emp_name, employees[i].dob, employees[i].role, employees[i].dept))
	}
	fmt.Println(db.Exec("select * from employees"))
	fmt.Println(employees)
}