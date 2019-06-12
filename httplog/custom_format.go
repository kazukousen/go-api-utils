package httplog

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

// https://httpd.apache.org/docs/2.2/logs.html#combined + execution time + requestID.
const apacheFormatPattern = "%s - - [%s] \"%s %s %s\" %d %d \"%s\" \"%s\" %.4f [%s]\n"

// CustomLogFormatter ...
func CustomLogFormatter(w io.Writer) middleware.LogFormatter {
	return &customLogFormatter{writer: w}
}

type customLogFormatter struct {
	writer io.Writer
}

func (f customLogFormatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	referer := r.Referer()
	if referer == "" {
		referer = "-"
	}

	userAgent := r.UserAgent()
	if userAgent == "" {
		userAgent = "-"
	}

	reqID := GetRequestID(r.Context())

	entry := &logEntry{
		writer:    f.writer,
		ip:        r.RemoteAddr,
		method:    r.Method,
		uri:       r.RequestURI,
		protocol:  r.Proto,
		referer:   referer,
		userAgent: userAgent,
		reqID:     reqID,
	}
	return entry
}

type logEntry struct {
	writer io.Writer

	ip                    string
	method, uri, protocol string
	referer, userAgent    string
	reqID                 string
}

func (l logEntry) Write(status, bytes int, elapsed time.Duration) {
	timeFormatted := time.Now().Format("02/Jan/2006 03:04:05")
	fmt.Fprintf(l.writer, apacheFormatPattern, l.ip, timeFormatted, l.method,
		l.uri, l.protocol, status, bytes, l.referer, l.userAgent, elapsed.Seconds(), l.reqID)
}

func (l logEntry) Panic(v interface{}, stack []byte) {
}
