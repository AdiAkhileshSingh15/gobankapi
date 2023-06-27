package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(fname, lname, pw string, store Storage) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account =>", acc.Number)

	return acc
}

func seedAccounts(s Storage) {
	seedAccount("eddy", "xing", "eddy123", s)
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("seeding the database")
		seedAccounts(store)
	}

	server := NewAPIServer(":8000", store)
	server.Run()
}
