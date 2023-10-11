package osv_test

import (
	"testing"

	"github.com/khulnasoft-lab/go-vulndb/internal/test"
)

func TestImports(t *testing.T) {
	// package osv only allows non stdlib imports.
	//
	// This is intended to make it easy for anyone to copy and paste the
	// JSON structs if needed.
	test.VerifyImports(t)
}
