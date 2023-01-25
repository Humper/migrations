package main

import (
	"fmt"

	"github.com/Humper/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("adding id column...")
		_, err := db.Exec(`ALTER TABLE my_table ADD id serial`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping id column...")
		_, err := db.Exec(`ALTER TABLE my_table DROP id`)
		return err
	})
}
