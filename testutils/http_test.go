package testutils_test

import (
	"net/http"
	"testing"

	"github.com/kazukousen/go-api-utils/testutils"
)

func TestRequestHTTP(t *testing.T) {
	tests := []struct {
		desc                  string
		handler               http.Handler
		method, path, payload string
		wantCode              int
		ct                    testutils.ContentType
		treePath, wantBody    string
	}{
		{desc: "JSON API",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(200)
				w.Write([]byte(`
				{
					"foo": [
						{"bar": "baz"},
						{"bar": "faz"}
					]
				}
				`))
			}), method: "GET", path: "/", payload: "",
			wantCode: 200, ct: testutils.JSON, treePath: "/foo[1]/bar", wantBody: "faz",
		},
		{desc: "XML API",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/xml")
				w.WriteHeader(200)
				w.Write([]byte(`
				<Foo>
					<Bar>baz</Bar>
					<Bar>faz</Bar>
				</Foo>
				`))
			}), method: "GET", path: "/", payload: "",
			wantCode: 200, ct: testutils.XML, treePath: "/Foo/Bar[2]", wantBody: "faz",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()
			if err := testutils.RequestHTTP(tt.handler, tt.method, tt.path, tt.payload, tt.wantCode, tt.ct, tt.treePath, tt.wantBody); err != nil {
				t.Errorf("error: %+v", err)
			}
		})
	}
}
