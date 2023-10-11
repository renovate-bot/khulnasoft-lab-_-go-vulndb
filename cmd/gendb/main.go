package main

import (
	"context"
	"flag"
	"log"

	db "github.com/khulnasoft-lab/go-vulndb/internal/database"
	"github.com/khulnasoft-lab/go-vulndb/internal/database/legacydb"
	"github.com/khulnasoft-lab/go-vulndb/internal/gitrepo"
)

var (
	repoDir = flag.String("repo", ".", "Directory containing vulndb repo")
	jsonDir = flag.String("out", "out", "Directory to write JSON database to")
	indent  = flag.Bool("indent", false, "Indent JSON for debugging")
	legacy  = flag.Bool("legacy", false, "if true, generate in the legacy schema")
)

func main() {
	flag.Parse()
	ctx := context.Background()
	repo, err := gitrepo.CloneOrOpen(ctx, *repoDir)
	if err != nil {
		log.Fatal(err)
	}
	if *legacy {
		if err := legacydb.Generate(ctx, repo, *jsonDir, *indent); err != nil {
			log.Fatal(err)
		}
	} else {
		d, err := db.FromRepo(ctx, repo)
		if err != nil {
			log.Fatal(err)
		}
		if err := d.Write(*jsonDir); err != nil {
			log.Fatal(err)
		}
	}
}
