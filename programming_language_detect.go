package text

import (
	enry "gopkg.in/src-d/enry.v1"
)

func DetectLanguage(text string) (string, bool) {
	lang, safe := enry.GetLanguageByContent(text)
	return lang, safe
}
