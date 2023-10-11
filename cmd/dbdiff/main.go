package main

import (
	"fmt"
	"log"
	"os"

	"github.com/khulnasoft-lab/go-vulndb/internal/database/legacydb"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: dbdiff db-a db-b")
		os.Exit(1)
	}
	if err := legacydb.Diff(os.Args[1], os.Args[2]); err != nil {
		log.Fatal(err)
	}
}
