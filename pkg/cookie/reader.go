// Package cookie
// parser netscape cookie file
// special thanks to https://unix.stackexchange.com/questions/36531/format-of-cookies-when-using-wget
package cookie

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	hdr = `# Netscape HTTP Cookie File`
)

// Import read netscape cookie file and return cookie jar
func Import(r io.Reader, u *url.URL) (http.CookieJar, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	cookies := []*http.Cookie{}
	for s.Scan() {
		txt := s.Text()
		if txt == hdr {
			continue
		}
		col := strings.Split(txt, "\t")
		if len(col) != 7 {
			continue
		}

		expire, err := strconv.Atoi(col[4])
		if err != nil {
			return nil, fmt.Errorf("failed to parse line for int: %v, column: 4: %v with err: %v", txt, col[4], err)
		}
		expiretime := time.Unix(int64(expire), 0)

		secure := false
		if col[3] == "TRUE" {
			secure = true
		}
		cookies = append(cookies,
			&http.Cookie{
				Name:    col[5],
				Value:   col[6],
				Path:    col[2],
				Domain:  col[0],
				Expires: expiretime,
				Secure:  secure,
			})
	}
	jar.SetCookies(u, cookies)
	return jar, nil
}

// ImportFromBase64 util to import cookies from base64 encoded netscape cookie file
func ImportFromBase64(e string, u *url.URL) (http.CookieJar, error) {
	data, err := base64.StdEncoding.DecodeString(e)

	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(data)
	return Import(r, u)
}
