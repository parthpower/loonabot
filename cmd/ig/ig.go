package main

import (
	"fmt"
	"os"

	"github.com/parthpower/loonabot/pkg/insta"

	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <instagram url>\n", os.Args[0])
		os.Exit(1)
	}
	godotenv.Load(".env")
	cookies := os.Getenv("INSTA_COOKIES")
	if cookies == "" {
		fmt.Fprintf(os.Stderr, "can't find cookies in INSTA_COOKIES")
		os.Exit(1)
	}
	media, err := insta.GetMediaFromUrl(os.Args[1], cookies)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Caption: %s\n", media.Caption)
	if media.Type == insta.TypeCarousel {
		for _, u := range media.MediaList {
			fmt.Println(u.URL)
		}
	} else {
		fmt.Println(media.MediaURL)
	}
}
