package main

import (
	"fmt"
	"os"

	"github.com/sagar/go-postgres-parser/internal/parse"
)

func main() {
	os.Exit(run())
}

func run() int {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "help" || args[0] == "-h" || args[0] == "--help" {
		usage()
		if len(args) == 0 {
			return 1
		}
		return 0
	}
	if args[0] == "version" {
		fmt.Println("pgp 0.1.0")
		return 0
	}

	raw := args[0]
	fields, err := parse.PostgresURL(raw)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("Host: %s\n", fields.Host)
	fmt.Printf("Port: %s\n", fields.Port)
	fmt.Printf("Database Name: %s\n", fields.Database)
	fmt.Printf("Username: %s\n", fields.Username)
	fmt.Printf("Password: %s\n", fields.Password)
	return 0
}

func usage() {
	fmt.Fprintf(os.Stderr, `pgp — parse a Postgres connection URL (local, no network)

Usage:
  pgp <postgres-url>

Prints Host, Port, Database Name, Username, and Password. Does not connect to the database.

Example:
  pgp 'postgres://neondb_owner:secret@ep-cool-darkness-123456.us-east-2.aws.neon.tech:5432/neondb?sslmode=require'

`)
}
