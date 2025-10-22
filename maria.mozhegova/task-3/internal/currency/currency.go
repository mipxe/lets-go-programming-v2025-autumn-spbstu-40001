package currency

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type ValCurs struct {
	Valutes []struct {
		NumCode  int           `xml:"NumCode"`
		CharCode string        `xml:"CharCode"`
		Value    CustomFloat64 `xml:"Value"`
	} `xml:"Valute"`
}

type CustomFloat64 float64

func (c *CustomFloat64) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var valueStr string
	err := decoder.DecodeElement(&valueStr, &start)
	if err != nil {
		return fmt.Errorf("Failed to parse value: %w", err)
	}

	value, err := strconv.ParseFloat(strings.Replace(valueStr, ",", ".", 1), 64)
	if err != nil {
		return fmt.Errorf("Failed to parse value: %w", err)
	}

	*c = CustomFloat64(value)
	return nil
}

func ReadValCurs(path string) (*ValCurs, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read xml file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = func(encoding string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, encoding)
	}

	var valCurs ValCurs
	err = decoder.Decode(&valCurs)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse XML: %w", err)
	}

	return &valCurs, nil
}
