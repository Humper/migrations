package main

import (
	"fmt"

	"github.com/Humper/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table my_table...")
		_, err := db.Exec(`CREATE TABLE my_table()`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table my_table...")
		_, err := db.Exec(`DROP TABLE my_table`)
		return err
	})
}
