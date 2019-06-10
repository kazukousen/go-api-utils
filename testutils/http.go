package testutils

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"golang.org/x/xerrors"
)

// RequestHTTP ...
func RequestHTTP(handler http.Handler, method, path, payload string, wantCode int, ct ContentType, treePath string, wantBody string) error {
	body, err := request(handler, method, path, payload, wantCode)
	if err != nil {
		return err
	}

	got, err := ScanBody(ct, body, treePath)
	if err != nil {
		return err
	}

	if got != wantBody {
		return xerrors.Errorf("not equal Body:\n\ngot:\n%s\nwant:\n%s\n", got, wantBody)
	}

	return nil
}

func request(handler http.Handler, method, path, payload string, wantCode int) (io.Reader, error) {
	srv := httptest.NewServer(handler)
	defer srv.Close()
	path = srv.URL + path

	req, err := http.NewRequest(method, path, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != wantCode {
		return nil, xerrors.Errorf("not equal StatusCode: got %d, but want %d", res.StatusCode, wantCode)
	}

	buf := new(bytes.Buffer)
	io.Copy(buf, res.Body)

	return buf, nil
}
