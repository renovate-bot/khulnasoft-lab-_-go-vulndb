package main

import (
	"flag"
	"log"

	db "github.com/khulnasoft-lab/go-vulndb/internal/database"
	"github.com/khulnasoft-lab/go-vulndb/internal/database/legacydb"
)

var (
	newPath       = flag.String("new", "", "path to new database")
	newLegacyPath = flag.String("legacy", "", "path to the new database in the legacy schema (optional)")
	existingPath  = flag.String("existing", "", "path to existing database")
)

func main() {
	flag.Parse()
	if *newPath == "" {
		log.Fatalf("flag -new must be set")
	}
	if *existingPath == "" {
		log.Fatalf("flag -existing must be set")
	}
	if err := db.Validate(*newPath, *existingPath); err != nil {
		log.Fatal(err)
	}
	if *newLegacyPath != "" {
		if err := legacydb.Validate(*newLegacyPath, *existingPath); err != nil {
			log.Fatal(err)
		}
		if err := legacydb.Equivalent(*newPath, *newLegacyPath); err != nil {
			log.Fatal(err)
		}
	}
}
