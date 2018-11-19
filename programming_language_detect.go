package text

import (
	enry "gopkg.in/src-d/enry.v1"
)

func DetectProgrammingLanguage(text string) (string, bool) {
	lang, safe := enry.GetLanguageByModeline([]byte(text))
	return lang, safe
}
