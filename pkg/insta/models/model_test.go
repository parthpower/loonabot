package models

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

var testjson []string

func init() {
	testjson = []string{
		"issue.json",
		"pic.json",
		"vid.json",
		"carouselpic.json",
		"carouselvid.json",
		"carouselpicvid.json",
		"igtv.json",
		"reel.json",
	}
}

func TestModel(t *testing.T) {
	for _, j := range testjson {
		t.Logf("parsing %s", j)
		_, err := parseFile(j)
		if err != nil {
			t.Fail()
			t.Logf("failed to parse: %s error: %v", j, err)
		}
	}
}

func parseFile(name string) (*Insta, error) {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	i := Insta{}
	err = json.Unmarshal(b, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}
