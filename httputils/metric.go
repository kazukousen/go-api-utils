package httputils

import (
	"net/http"
	"strconv"
	"time"

	"github.com/kazukousen/go-api-utils/metrics"
)

func metric(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := &wrapResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		start := time.Now()
		defer func() {
			metrics.HTTPRequestLatency.
				WithLabelValues(r.Method, strconv.Itoa(ww.statusCode)).
				Observe(time.Since(start).Seconds() * 1000)
		}()
		next.ServeHTTP(ww, r)
	})
}

type wrapResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (ww *wrapResponseWriter) WriteHeader(statusCode int) {
	ww.statusCode = statusCode
	ww.ResponseWriter.WriteHeader(statusCode)
}
