package main

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "050220"
	dbname   = "test"
)

func main() {

	psqlcomm := fmt.Sprintf("host %s port %d user is %s password %s dbname %s sslmode = disabled", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlcomm)
	CheckError(err)

	defer db.Close()
	insertStmt := `insert into  "Employee("name","EmpId") values("Rohit",21)`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	insertDynstmt := `insert into "Employee"("Name","EmpId") values($1,$2)`
	_, e = db.Exec(insertDynstmt, "krishna", 01)
	CheckError(e)

}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
