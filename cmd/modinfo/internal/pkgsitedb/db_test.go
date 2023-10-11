package pkgsitedb

import (

	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var dbConfigFile = flag.String("dbconf", "", "filename with db config as JSON")

func TestQueryModule(t *testing.T) {
	if *dbConfigFile == "" {
		t.Skip("missing -dbconf")
	}
	f, err := os.Open(*dbConfigFile)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	var cfg Config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	db, err := Open(ctx, cfg)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if err := db.PingContext(ctx); err != nil {
		t.Fatal(err)
	}
	m, err := QueryModule(ctx, db, "golang.org/x/tools")
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range m.Packages {
		fmt.Printf("%+v\n", p)
	}
}
