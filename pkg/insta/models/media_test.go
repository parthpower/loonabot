package models

import "testing"

func TestDownloadURL(t *testing.T) {
	for _, j := range testjson {
		t.Logf("fetching for %s", j)
		i, err := parseFile(j)
		if err != nil {
			t.Logf("failed to parse %s error %v", j, err)
			t.Fail()
		}
		urls, err := i.DownloadURL()
		if err != nil {
			t.Logf("failed to get download url for %s error %v", j, err)
			t.Fail()
		}
		if len(urls) == 0 {
			t.Logf("urls not populated for %s", j)
			t.Fail()
		} else {
			t.Logf("%s: %v", j, urls)
		}
	}
}
