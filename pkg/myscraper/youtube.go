package myscraper

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/buger/jsonparser"
)

type SearchResult struct {
	URL string
}

func getJsonPayload() []byte {
	b := `{
		"context": {
			  "client": {
				  "clientName": "WEB",
				  "clientVersion": "2.20210224.06.00",
				  "newVisitorCookie": "true"
			  },
			  "user": {
				  "lockedSafetyMode": "false"
			  }
		  },
		 "query":"",
		 "client": {"hl":"en","gl":"US"},
		 "params": "EgIQAQ%3D%3D"
	  }`
	return []byte(b)
}

const useragent = `User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36`
const url = `https://www.youtube.com/youtubei/v1/search?key=AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8`
const contentType = `application/json; charset=utf-8`

func Search(term string) ([]SearchResult, error) {

	body, err := jsonparser.Set(getJsonPayload(), []byte(fmt.Sprintf(`"%s"`, term)), "query")
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(body)
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", useragent)
	req.Header.Set("Content-Type", contentType)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	l := []SearchResult{}
	_, err = jsonparser.ArrayEach(r, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if err != nil {
			return
		}
		vid, err := jsonparser.GetString(value, "videoRenderer", "videoId")
		if err == nil {
			l = append(l, SearchResult{URL: "https://youtube.com/watch?v=" + vid})
		}
	}, "contents", "twoColumnSearchResultsRenderer", "primaryContents", "sectionListRenderer", "contents", "[0]", "itemSectionRenderer", "contents")
	// vidId, err := jsonparser.GetString(r,"contents","twoColumnSearchResultsRenderer","primaryContents","sectionListRenderer","contents","[0]","itemSectionRenderer","contents","[0]","videoRenderer","videoId")
	if err != nil {
		return nil, err
	}
	return l, err
}
