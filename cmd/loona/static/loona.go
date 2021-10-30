package static

import (
	"embed"
	"io/fs"
)

//go:embed *.mp4
var f embed.FS

func WelcomeToLoona() (fs.File, error) {
	return File("welcometoloona.mp4")
}

func KimLip() (fs.File, error) {
	return File("kimlip.mp4")
}

func ImCat() (fs.File, error) {
	return File("imcat.mp4")
}

func Haseul() (fs.File, error) {
	return File("haseul.mp4")
}

func File(name string) (fs.File, error) {
	return f.Open(name)
}
