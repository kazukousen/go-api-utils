package httputils_test

import (
	"strings"
	"testing"

	"github.com/kazukousen/go-api-utils/httputils"
)

func TestScanBody(t *testing.T) {
	tests := []struct {
		desc          string
		ct            httputils.ContentType
		body, p, want string
	}{
		{desc: "scanJSON", ct: httputils.JSON, body: `
		{
			"foo": [
				{"bar": "baz"},
				{"bar": "faz"}
			]
		}
		`, p: "/foo[1]/bar", want: "faz"},
		{desc: "scanXML", ct: httputils.XML, body: `
		<Foo>
			<Bar>baz</Bar>
			<Bar>faz</Bar>
		</Foo>
		`, p: "/Foo/Bar[2]", want: "faz"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()
			got, err := httputils.ScanBody(tt.ct, strings.NewReader(tt.body), tt.p)
			if err != nil {
				t.Errorf("could not scan body: %+v", err)
			}
			if got != tt.want {
				t.Errorf("got:\n%s\nbut want:\n%s\n", got, tt.want)
			}
		})
	}
}
