package report

import (
	"strings"

	"github.com/khulnasoft-lab/go-vulndb/internal/osv"
)

func isFix(url string) bool {
	return strings.Contains(url, "go-review.googlesource.com") ||
		strings.Contains(url, "/commit/") || strings.Contains(url, "/commits/") ||
		strings.Contains(url, "/pull/") || strings.Contains(url, "/cl/")
}

func isIssue(url string) bool {
	return strings.Contains(url, "/issue/") || strings.Contains(url, "/issues/")
}

func isAdvisory(url string) bool {
	return strings.Contains(url, "/advisories/")
}

// referenceFromUrl creates a new Reference from a url
// with Type inferred from the contents of the url.
func referenceFromUrl(url string) *Reference {
	typ := osv.ReferenceTypeWeb
	switch {
	case isFix(url):
		typ = osv.ReferenceTypeFix
	case isIssue(url):
		typ = osv.ReferenceTypeReport
	case isAdvisory(url):
		typ = osv.ReferenceTypeAdvisory
	}
	return &Reference{
		Type: typ,
		URL:  url,
	}
}
