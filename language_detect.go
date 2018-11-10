package text

import "github.com/abadojack/whatlanggo"

func DetectLanguage(text string) (string, error) {
	info := whatlanggo.Detect(text)
	return whatlanggo.LangToString(info.Lang.Str), nil
}
