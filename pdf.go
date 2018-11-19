package text

import (
	"bytes"

	pdfcontent "github.com/unidoc/unidoc/pdf/contentstream"
	pdf "github.com/unidoc/unidoc/pdf/model"
)

func ReadPDF(bts0 []byte) (string, error) {
	bts := bytes.NewReader(bts0)
	res := &bytes.Buffer{}

	pdfReader, err := pdf.NewPdfReader(bts)
	if err != nil {
		return "", err
	}

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return "", err
	}

	if isEncrypted {
		_, err = pdfReader.Decrypt([]byte(""))
		if err != nil {
			return "", err
		}
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return "", err
	}

	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return "", err
		}

		contentStreams, err := page.GetContentStreams()
		if err != nil {
			return "", err
		}

		// If the value is an array, the effect shall be as if all of the streams in the array were concatenated,
		// in order, to form a single stream.
		pageContentStr := ""
		for _, cstream := range contentStreams {
			pageContentStr += cstream
		}

		cstreamParser := pdfcontent.NewContentStreamParser(pageContentStr)
		txt, err := cstreamParser.ExtractText()
		if err != nil {
			return "", err
		}
		res.Write([]byte(txt))
	}

	return res.String(), nil
}
