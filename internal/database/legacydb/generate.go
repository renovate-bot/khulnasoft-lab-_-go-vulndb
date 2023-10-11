package legacydb

import (
	"context"

	"github.com/go-git/go-git/v5"
	"github.com/khulnasoft-lab/go-vulndb/internal/derrors"
)

// Generate creates and writes a new Go vulnerability database to outDir
// based on the contents of the "data/osv" folder inside repo.
//
// The repo must contain a "data/osv" folder with files in
// OSV JSON format with filenames of the form GO-YYYY-XXXX.json.
func Generate(ctx context.Context, repo *git.Repository, outDir string, indent bool) (err error) {
	defer derrors.Wrap(&err, "Generate()")

	new, err := New(ctx, repo)
	if err != nil {
		return err
	}
	if err = new.Write(outDir, indent); err != nil {
		return err
	}

	return nil
}
