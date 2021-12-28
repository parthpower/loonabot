package cookie

import (
	"encoding/base64"
	"net/url"
	"strings"
	"testing"
)

const cookie = `# Netscape HTTP Cookie File

.instagram.com	TRUE	/	TRUE	1703561918	mid	lajsdlajsd
.instagram.com	TRUE	/	TRUE	1703561918	ig_did	asdasdasdasd
.instagram.com	TRUE	/	TRUE	1672025918	ig_nrcb	1234
.instagram.com	TRUE	/	TRUE	1672026436	sessionid	ljlskajd%3A14
.instagram.com	TRUE	/	TRUE	1671995627	csrftoken	dd3423
.instagram.com	TRUE	/	TRUE	0	rur	"jjjjjjjjjjjjjjj"
.instagram.com	TRUE	/	TRUE	1648322027	ds_user_id	1232455123123
`

func TestImport(t *testing.T) {
	r := strings.NewReader(cookie)
	u, _ := url.Parse("https://instagram.com")
	jar, err := Import(r, u)
	if err != nil {
		t.Fail()
		t.Log(err)
	}
	t.Log(jar)
}

func TestImportBase64(t *testing.T) {
	e := base64.StdEncoding.EncodeToString([]byte(cookie))
	u, _ := url.Parse("https://instagram.com")
	jar, err := ImportFromBase64(e, u)
	if err != nil {
		t.Fail()
		t.Log(err)
	}
	t.Log(jar)
}
