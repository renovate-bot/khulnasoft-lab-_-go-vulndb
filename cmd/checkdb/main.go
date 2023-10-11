package main

import (
	"flag"
	"log"

	db "github.com/khulnasoft-lab/go-vulndb/internal/database"
	"github.com/khulnasoft-lab/go-vulndb/internal/database/legacydb"
)

var legacy = flag.Bool("legacy", false, "if true, check with respect to legacy database schema")

func main() {
	flag.Parse()
	path := flag.Arg(0)
	if path == "" {
		log.Fatal("path must be set\nusage: checkdb [path]")
	}
	if *legacy {
		if _, err := legacydb.Load(path); err != nil {
			log.Fatal(err)
		}
	} else {
		if _, err := db.Load(path); err != nil {
			log.Fatal(err)
		}
	}
}
