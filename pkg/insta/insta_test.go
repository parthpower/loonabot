package insta

import (
	"os"
	"regexp"
	"testing"
)

func TestMedia(t *testing.T) {
	cookies := os.Getenv("INSTA_COOKIES")
	if cookies == "" {
		t.Log("Export INSTA_COOKIES env var to run this test")
		t.SkipNow()
	}
	m, err := GetPostFromCode("CW8BdQlhono", cookies)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(m.MediaURL)
}

func TestReel(t *testing.T) {
	cookies := os.Getenv("INSTA_COOKIES")
	if cookies == "" {
		t.Log("Export INSTA_COOKIES env var to run this test")
		t.SkipNow()
	}
	m, err := GetReelFromCode("CX5-n1ylthH", cookies)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(m.MediaURL)
	t.Log(m.Type)
}

func TestURL(t *testing.T) {
	cookies := os.Getenv("INSTA_COOKIES")
	if cookies == "" {
		t.Log("Export INSTA_COOKIES env var to run this test")
		t.SkipNow()
	}
	m, err := GetMediaFromUrl("https://www.instagram.com/tv/CYALWwdJ54u/?utm_medium=copy_link", cookies)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(m.MediaURL)
	t.Log(m.Type)
}

func TestRegex(t *testing.T) {
	testString := `https://www.instagram.com/tv/CYALWwdJ54u/?utm_medium=copy_link
	https://www.instagram.com/loonathedorm/reel/CX_mcNRAbx6/
	`
	re := regexp.MustCompile(`https?:\/\/(?:www\.)?instagram\.com\/(:?.*\/)?(:?tv|p|reel)\/([A-Za-z0-9_-]+)`)

	matches := re.FindAllStringSubmatch(testString, -1)
	if len(matches) < 1 {
		t.FailNow()
	}
	for _, match := range matches {
		if len(match) > 1 {
			t.Log(match)
		}
	}
}
