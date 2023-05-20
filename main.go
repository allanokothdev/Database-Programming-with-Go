package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//connect to the database
	db, err := sql.Open("mysql", "root:allanokothdev@tcp(127.0.0.1:3306)/mytestdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//creating table through SQL DDL Commands
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS mytestdb.emp(id int, name varchar(50))")
	if err != nil {
		panic(err.Error())
	}

	//Inserting 5 records
	result, err := db.Exec("INSERT INTO mytestdb.emp(id,name,email) VALUES (101, 'mickey', 'mickey@gmail.com),(102, 'Donald','donald@gmail.com),(103,'Tom','tom@gmail.com'),(104,'Jerry','jerry@gmail.com'),(105,'Minnie','minnie@gmail.com')")

	if err != nil {
		panic(err.Error())
	}

	rc, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Inserted %d rows\n", rc)


	//Simple select query to fetch all records
	rows, err := db.Query("SELECT * FROM mytestdb.emp")
	if err != nil {
		panic(err.Error())
	}

	//	Struct
	type emp struct {
		ID int
		Name string
		Email string
	}

	for rows.Next() {
		var e emp
		err = rows.Scan(&e.ID, &e.Name, &e.Email)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%d %s %s \n", e.ID, e.Name, e.Email)
	}
}