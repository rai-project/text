package text

import slugify "github.com/mozillazg/go-slugify"

func Slugify(text string) (string, error) {
	return slugify.Slugify(text), nil
}
