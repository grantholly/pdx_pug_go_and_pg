package main

import (
	"fmt"

	"github.com/go-pg/pg"
)

type Item struct {
	Id    int
	Items []Item `pg:",many2many:item_to_items,joinFK:Sub"`
}

func main() {

	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "demo",
		PoolSize: 10,
		Addr:     "192.168.0.100:5432",
	})

	defer db.Close()

	qs := []string{
		"CREATE TEMP TABLE items (id int)",
		"CREATE TEMP TABLE item_to_items (item_id int, sub_id int)",
		"INSERT INTO items VALUES (1), (2), (3)",
		"INSERT INTO item_to_items VALUES (1, 2), (1, 3)",
	}
	for _, q := range qs {
		_, err := db.Exec(q)
		if err != nil {
			panic(err)
		}
	}

	// Select item and all subitems with following queries:
	//
	// SELECT "item".* FROM "items" AS "item" ORDER BY "item"."id" LIMIT 1
	//
	// SELECT * FROM "items" AS "item"
	// JOIN "item_to_items" ON ("item_to_items"."item_id") IN ((1))
	// WHERE ("item"."id" = "item_to_items"."sub_id")

	var item Item
	err := db.Model(&item).Column("item.*", "Items").First()
	if err != nil {
		panic(err)
	}
	fmt.Println("Item", item.Id)
	fmt.Println("Subitems", item.Items[0].Id, item.Items[1].Id)
}
