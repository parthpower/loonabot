package insta

import (
	"os"
	"regexp"
	"testing"
)

func TestURL(t *testing.T) {
	cookies := os.Getenv("INSTA_COOKIES")
	if cookies == "" {
		t.Log("Export INSTA_COOKIES env var to run this test")
		t.SkipNow()
	}
	urls := []string{
		"https://www.instagram.com/loonathedorm/reel/CX_mcNRAbx6/",
		"https://www.instagram.com/tv/CYALWwdJ54u/?utm_medium=copy_link",
		"https://www.instagram.com/p/CY1FWZ0J-ew/",
		"https://www.instagram.com/p/CYLi2doLHVE/",
	}
	for _, url := range urls {
		m, err := GetMediaFromUrl(url, cookies)
		if err != nil {
			t.Logf("failed to get url %s", url)
			t.Log(err)
			t.Fail()
		}
		t.Log(m.Caption)
		t.Log(m.DownloadURL)
	}
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
