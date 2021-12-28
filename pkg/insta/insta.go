// Package insta
// something to pull instagram resources
package insta

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/parthpower/loonabot/pkg/cookie"
)

const (
	PostEndpoint = "https://instagram.com/p"
	ReelEndpoint = "https://instagram.com/p"
)

func GetPostFromCode(code string, cookies string) (Media, error) {
	return getMediaFromCode(code, cookies, PostEndpoint)
}

func GetReelFromCode(code string, cookies string) (Media, error) {
	return getMediaFromCode(code, cookies, ReelEndpoint)
}

func GetMediaFromUrl(u string, cookies string) (Media, error) {
	instaurl, _ := url.Parse("https://instagram.com")
	jar, err := cookie.ImportFromBase64(cookies, instaurl)
	if err != nil {
		return Media{}, err
	}
	c := &http.Client{
		Transport: http.DefaultTransport,
		Jar:       jar,
	}
	uu, err := url.Parse(u)
	if err != nil {
		return Media{}, err
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
		return Media{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Media{}, err
	}
	return getFromMediaPage(body)

}

func getMediaFromCode(code string, cookies string, endpoint string) (Media, error) {
	instaurl, _ := url.Parse("https://instagram.com")
	jar, err := cookie.ImportFromBase64(cookies, instaurl)
	if err != nil {
		return Media{}, err
	}
	c := &http.Client{
		Transport: http.DefaultTransport,
		Jar:       jar,
	}
	u, _ := url.Parse(fmt.Sprintf("%s/%s/?__a=1", endpoint, code))
	resp, err := c.Do(&http.Request{
		Method: http.MethodGet,
		URL:    u,
		Proto:  "HTTP/2",
		Header: http.Header{
			"Connection": []string{"keep-alive"},
		},
	})
	if err != nil {
		return Media{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Media{}, err
	}
	return getFromMediaPage(body)
}
