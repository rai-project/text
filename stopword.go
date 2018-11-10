package text

import "github.com/bbalet/stopwords"

func RemoveStopWords(text string) (string, error) {
	lang, err := DetectLanguage(text)
	if err != nil {
		return "", err
	}
	return stopwords.CleanString(text, lang, true)
}
