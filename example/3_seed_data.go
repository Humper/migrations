package main

import (
	"fmt"

	"github.com/Humper/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("seeding my_table...")
		_, err := db.Exec(`INSERT INTO my_table VALUES (1)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("truncating my_table...")
		_, err := db.Exec(`TRUNCATE my_table`)
		return err
	})
}
