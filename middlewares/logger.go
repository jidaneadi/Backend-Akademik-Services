package middlewares

import (
	"log"
	"net/http"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(lrw, r)

		logger.Printf(
			"%s - [%s] \"%s %s %s\" %d %s",
			r.RemoteAddr,
			start.Format(time.RFC1123),
			r.Method,
			r.RequestURI,
			r.Proto,
			lrw.statusCode,
			time.Since(start),
		)
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
