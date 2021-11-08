package myscraper

import (
	"testing"
)

func TestSearch(t *testing.T) {
	r, err := Search("hula hoop loona")
	if err != nil {
		t.Fail()
	}
	if len(r) < 1 {
		t.Fail()
	}
	t.Log(r[0].URL)
}
