package httplog

import "github.com/go-chi/chi/middleware"

// GetRequestID ...
var GetRequestID = middleware.GetReqID
