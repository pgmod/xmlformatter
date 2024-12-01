package xmlformatter

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

func FormatXML(xmlData string) (string, error) {

	buffer := bytes.NewBuffer([]byte(xmlData))
	decoder := xml.NewDecoder(buffer)
	var output bytes.Buffer
	encoder := xml.NewEncoder(&output)
	encoder.Indent("", "\t")

	for {
		tok, err := decoder.Token()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return "", fmt.Errorf("error decoding XML: %w", err)
		}
		if err := encoder.EncodeToken(tok); err != nil {
			return "", fmt.Errorf("error encoding XML: %w", err)
		}
	}

	if err := encoder.Flush(); err != nil {
		return "", fmt.Errorf("error flushing XML: %w", err)
	}

	return output.String(), nil
}
