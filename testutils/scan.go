package testutils

import (
	"io"
	"strings"

	"github.com/mattn/go-scan"
	"golang.org/x/xerrors"
	xmlpath "gopkg.in/xmlpath.v2"
)

// ContentType ...
type ContentType string

// content-type
const (
	JSON ContentType = "application/json"
	XML  ContentType = "application/xml"
)

// ScanBody ...
func ScanBody(ct ContentType, body io.Reader, p string) (got string, err error) {
	switch ct {
	case JSON:
		got, err = scanJSON(body, p)
	case XML:
		got, err = scanXML(body, p)
	}
	got = strings.TrimRight(got, "\n")
	return
}

func scanJSON(body io.Reader, p string) (string, error) {
	var got string
	if err := scan.ScanJSON(body, p, &got); err != nil {
		return "", err
	}
	return got, nil
}

func scanXML(body io.Reader, p string) (string, error) {
	xpath, err := xmlpath.Compile(p)
	if err != nil {
		return "", err
	}
	node, err := xmlpath.Parse(body)
	if err != nil {
		return "", err
	}
	got, ok := xpath.String(node)
	if !ok {
		return "", xerrors.Errorf("not found node: %s", p)
	}
	return got, nil
}
