package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

func main() {

    db, err := sql.Open("mysql", "golang:golang@tcp(127.0.0.1:3306)/testdb")
    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }

    sql := "CREATE TABLE IF NOT EXISTS person (id INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY, name VARCHAR(100));"
    res, err := db.Exec(sql)

    if err != nil {
        panic(err.Error())
    }

		i := 0

		for {
			sql = fmt.Sprintf("INSERT INTO person(name) VALUES ('%08d')", i)
			res, err = db.Exec(sql)
	
			if err != nil {
					panic(err.Error())
			}
	
			affectedRows, err := res.RowsAffected()
	
			if err != nil {
					log.Fatal(err)
			}
	
			fmt.Printf("inserted %08d, affected: %d\n", i, affectedRows)
			i += 1
		}

}