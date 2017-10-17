package main

import (
	"fmt"
	
	"github.com/jackc/pgx"
)

func connect(dsn string, maxConn int) *pgx.ConnPool {
     pgConfig, err := pgx.ParseURI(dsn)
     if err != nil {
     	fmt.Printf("cannot parse dsn %s", dsn)
     }

     conn, err := pgx.NewConnPool(pgx.ConnPoolConfig{
     	   ConnConfig: pgConfig,
	   MaxConnections: maxConn,
     })
     if err != nil {
     	fmt.Printf("cannot create connection pool")
     } else {
        fmt.Println("sweet!  we're in!")
     }

     return conn
}

func main() {

     dsn := "postgres://postgres:postgres@192.168.0.100:5432/demo"
     pool_size := 10

     db := connect(dsn, pool_size)
     
     fmt.Println(db)

     var id int
     var val string

     db.QueryRow("SELECT id, val from public.test where id = 2;").Scan(&id, &val)

     fmt.Println(id, val)
     
}
