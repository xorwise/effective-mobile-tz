package api

import (
	"log/slog"
	"net/http"
	"time"
)

type LoggingMiddleware struct {
	Logger *slog.Logger
}

func (lm LoggingMiddleware) WithLogging(next http.HandlerFunc) http.Handler {
	logFn := func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		start := time.Now()

		uri := r.RequestURI
		method := r.Method
		next(rw, r)

		duration := time.Since(start)

		lm.Logger.Info("request", "uri", uri, "method", method, "duration", duration)
	}
	return http.HandlerFunc(logFn)
}
