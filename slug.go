package text

func Slugify(text string) (string, error ) {
  return slugify.Slugify(text), nil
}
