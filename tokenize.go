package text

import prose "gopkg.in/jdkato/prose.v2"

func Tokenize(text string) ([]prose.Token, error) {
	doc, err := prose.NewDocument(text)
	if err != nil {
		return nil, err
	}
	return doc.Tokens(), nil
}
