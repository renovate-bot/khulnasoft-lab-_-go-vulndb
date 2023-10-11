//go:build go1.17
// +build go1.17

package store

import "testing"

func TestMemStore(t *testing.T) {
	testStore(t, NewMemStore())
}
