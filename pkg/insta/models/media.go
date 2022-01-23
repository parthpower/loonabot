package models

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

type Media struct {
	Caption     string
	DownloadURL []string
}

func GetInsta(r io.Reader) (*Insta, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var i Insta
	err = json.Unmarshal(b, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (i *Insta) Media() (*Media, error) {
	if i == nil {
		return nil, fmt.Errorf("insta is empty")
	}
	if len(i.Items) == 0 {
		return nil, fmt.Errorf("no items in insta")
	}

	u, err := i.DownloadURL()
	if err != nil {
		return nil, err
	}
	return &Media{
		Caption:     i.Items[0].Caption.Text,
		DownloadURL: u,
	}, nil
}

func (i *Insta) DownloadURL() ([]string, error) {
	if i == nil {
		return nil, fmt.Errorf("insta nil")
	}
	urls := []string{}
	for _, item := range i.Items {
		// for single picture
		u, _ := getImageCandidate(&item.ImageVersions2, item.OriginalWidth, item.OriginalHeight)
		// error is fine because post may not have image
		if u != "" {
			urls = append(urls, u)
		}
		vidu, _ := getVideoCandidate(&item.VideoVersions, item.OriginalWidth, item.OriginalHeight)
		if vidu != "" {
			urls = append(urls, vidu)
		}
		// check cart media
		for _, caritems := range item.CarouselMedia {
			u, _ := getImageCandidate(&caritems.ImageVersions2, caritems.OriginalWidth, caritems.OriginalWidth)
			// error is fine because post may not have image
			if u != "" {
				urls = append(urls, u)
			}

			// get videos
			vidu, _ := getVideoCandidate(&caritems.VideoVersions, caritems.OriginalWidth, caritems.OriginalHeight)
			if vidu != "" {
				urls = append(urls, vidu)
			}

		}
	}
	return urls, nil
}

// getImageCandidate get url for width, height. Can get these from OriginalWidth, OriginalHeight
func getImageCandidate(img *ImageVersions2, width, height int) (string, error) {
	for _, c := range img.Candidates {
		if c.Width == width && c.Height == height {
			return c.URL, nil
		}
	}
	return "", fmt.Errorf("image not found with %dx%d", width, height)
}

func getVideoCandidate(vid *[]VideoVersions, width, height int) (string, error) {
	for _, c := range *vid {
		if c.Width == width && c.Height == height {
			return c.URL, nil
		}
	}
	return "", fmt.Errorf("video not found with %dx%d", width, height)
}
