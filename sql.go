package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_HOST = "192.168.0.100"
	DB_PORT = "5432"
	DB_USER = "postgres"
	DB_PASS = "postgres"
	DB_NAME = "demo"
)

func BasicQuery(db sql.DB) {

	q_id := 1
	rows, err := db.Query("SELECT * from public.test where id > $1", q_id)
	if err != nil {
		fmt.Println("couldn't run the query")
	}

	defer rows.Close()

	for rows.Next() {
		var (
			val string
			id  int
		)
		err := rows.Scan(&id, &val)
		if err != nil {
			fmt.Println("couldn't assign values from query results")
		}
		fmt.Println(id, val)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println("got an error with the record set")
	}

}

func BasicSingleInsert(db sql.DB, id int, val string) {
	// Exec for statements that don't return a record set
	stmt, err := db.Prepare("INSERT INTO public.test (id, val) VALUES ($1, $2)")
	if err != nil {
		fmt.Println("couldn't create prepared statement")
	}

	// db.Exec("delete from test") won't burn a connection
	// db.Query("delete from test") burns a connection until sql.Rows.Close() is called
	res, err := stmt.Exec(id, val)
	if err != nil {
		fmt.Println("couldn't execute prepared statement")
	}

	fmt.Println("inserted", id, val)
	fmt.Println(res)
}

func BasicTransaction(db sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("couldn't open a transaction")
	}

	fmt.Println(tx)
	defer tx.Rollback()

	var cnt int
	err = tx.QueryRow("SELECT count(*) FROM public.test").Scan(&cnt)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cnt)

	if cnt > 5 {
		fmt.Sprintf("found %d rows", cnt)
	} else {
		fmt.Println("not enough rows found!  Performing an inserts")
		stmt, err := tx.Prepare("INSERT INTO public.test (id, val) VALUES ($1, $2)")
		if err != nil {
			fmt.Println("couldn't prepare statement inside transaction")
		}
		for i := 101; i < 111; i++ {
			res, err := stmt.Exec(i, "garbage row")
			if err != nil {
				fmt.Println("couldn't insert garbage row")
			}
			fmt.Println(res)
		}
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("couldn't commit transaction!  Rolling back")
	}
}

func main() {
	dbinfo := fmt.Sprintf("host=%s port=%s "+
		"user=%s password=%s "+
		"dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("could not open a connection to the DB")
	} else {
		fmt.Println("we're in!")
	}

	defer db.Close()

	fmt.Println(db)

	err = db.Ping()
	if err != nil {
		fmt.Println("couldn't ping the DB")
	}

	BasicQuery(*db)
	BasicSingleInsert(*db, 123, "whatever man")
	BasicTransaction(*db)
}
