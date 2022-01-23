// Package insta
// something to pull instagram resources
package insta

import (
	"net/http"
	"net/url"

	"github.com/parthpower/loonabot/pkg/cookie"
	"github.com/parthpower/loonabot/pkg/insta/models"
)

const (
	PostEndpoint = "https://instagram.com/p"
	ReelEndpoint = "https://instagram.com/reel"
)

func GetMediaFromUrl(u string, cookies string) (*models.Media, error) {
	instaurl, _ := url.Parse("https://instagram.com")
	jar, err := cookie.ImportFromBase64(cookies, instaurl)
	if err != nil {
		return nil, err
	}
	c := &http.Client{
		Transport: http.DefaultTransport,
		Jar:       jar,
	}
	uu, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	q, _ := url.ParseQuery("__a=1")
	uu.RawQuery = q.Encode()
	resp, err := c.Do(&http.Request{
		Method: http.MethodGet,
		URL:    uu,
		Proto:  "HTTP/2",
		Header: http.Header{
			"Connection": []string{"keep-alive"},
		},
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	i, err := models.GetInsta(resp.Body)
	if err != nil {
		return nil, err
	}

	return i.Media()

}
