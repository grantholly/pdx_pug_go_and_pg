package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

// github.com/jackc/pgx/stdlib is the driver version
func connect(dsn string, maxConn int) *pgx.ConnPool {
	pgConfig, err := pgx.ParseURI(dsn)
	if err != nil {
		fmt.Printf("cannot parse dsn %s", dsn)
	}

	conn, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     pgConfig,
		MaxConnections: maxConn,
	})
	if err != nil {
		fmt.Printf("cannot create connection pool")
	} else {
		fmt.Println("sweet!  we're in!")
	}

	return conn
}

func SelectQuery(db pgx.ConnPool) {
	var (
		id  int
		val string
	)

	row := db.QueryRow("SELECT id, val from public.test where id = 2;").Scan(&id, &val)

	fmt.Println(id, val)
	fmt.Println(row)
}

func InsertJson(db pgx.ConnPool, id int, doc string) {
	stmt, err := db.Prepare("json_insert",
		"INSERT INTO public.jsonb_test (id, doc) VALUES ($1, $2)")
	if err != nil {
		fmt.Println("couldn't prepare statement")
	}

	res, err := db.Exec(stmt.SQL, id, doc)
	if err != nil {
		fmt.Println("couldn't execute prepared insert statement")
	}

	fmt.Println("inserted", id, doc)
	fmt.Println(res)
}

func SelectJsonDoc(db pgx.ConnPool, id int) string {
	var doc string
	row := db.QueryRow("SELECT doc FROM public.jsonb_test WHERE id = $1", id)

	row.Scan(&doc)

	return doc
}

type Profile struct {
	Id      int                    `json:"id"`
	Name    string                 `json:"profileName"`
	Hobbies []string               `json:"hobbies"`
	Loc     map[string]interface{} `json:"location"`
}

func (p *Profile) ToJson() []byte {
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("couldn't serialze to JSON")
	}
	return b
}

func Deserialize(json_doc string) Profile {
	raw := []byte(json_doc)

	var p Profile

	err := json.Unmarshal(raw, &p)
	if err != nil {
		panic(err)
	}
	return p
}

func main() {

	dsn := "postgres://postgres:postgres@192.168.0.100:5432/demo"
	pool_size := 10

	db := connect(dsn, pool_size)

	fmt.Println(db)

	//SelectQuery(*db)

	json_doc := `
     	      {"id": 1,
	      "profileName": "admin",
	      "hobbies":["the great indoors"],
	      "location": {"state": "WA", "zip": 98632}}`

	InsertJson(*db, 1, json_doc)

	json_response := SelectJsonDoc(*db, 1)

	profile := Deserialize(json_response)
	os.Stdout.Write(profile.ToJson())
}
