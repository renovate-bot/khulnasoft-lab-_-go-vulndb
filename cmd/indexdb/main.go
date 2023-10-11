package main

import (
	"flag"
	"log"

	"github.com/khulnasoft-lab/go-vulndb/internal/database"
)

var (
	vulnsDir = flag.String("vulns", "", "Directory containing JSON OSV files")
	outDir   = flag.String("out", "", "Directory to write database to")
)

func main() {
	flag.Parse()
	if *vulnsDir == "" {
		log.Fatal("flag -vulns must be set")
	}
	if *outDir == "" {
		log.Fatal("flag -out must be set")
	}
	db, err := database.RawLoad(*vulnsDir)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Write(*outDir); err != nil {
		log.Fatal(err)
	}
}
