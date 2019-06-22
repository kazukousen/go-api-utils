package httputils

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/handlers"
	"github.com/kazukousen/go-api-utils/httplog"
)

// Middlewares ...
var Middlewares = []func(http.Handler) http.Handler{
	metric,
	middleware.RequestID,
	middleware.RequestLogger(httplog.CustomLogFormatter(os.Stdout)),
	middleware.Timeout(10 * time.Second),
	middleware.RealIP,
	handlers.CompressHandler,
	middleware.Recoverer,
}
