package worker

import (
	"encoding/json"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/khulnasoft-lab/go-vulndb/internal/cveschema"
	"github.com/khulnasoft-lab/go-vulndb/internal/derrors"
)

// ReadCVEAtPath reads file at path in commit, and JSON-decodes it into a CVE.
func ReadCVEAtPath(commit *object.Commit, path string) (_ *cveschema.CVE, blobHash string, err error) {
	defer derrors.Wrap(&err, "readCVEAtPath(%q)", path)
	file, err := commit.File(path)
	if err != nil {
		return nil, "", err
	}
	var cve cveschema.CVE
	r, err := file.Reader()
	if err != nil {
		return nil, "", err
	}
	if err := json.NewDecoder(r).Decode(&cve); err != nil {
		return nil, "", err
	}
	return &cve, file.Hash.String(), nil
}
