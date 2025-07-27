package middleware

import (
	"log"
	"net/http"
	"os"
	"time"
)

var logger *log.Logger

func init() {
	_ = os.MkdirAll("logs", os.ModePerm)
	logFile, err := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("error", err)
	}
	logger = log.New(logFile, ":", log.LstdFlags)
}
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: 200}
		next.ServeHTTP(rec, r)
		duration := time.Since(start)
		logger.Printf(
			"%s - %s %s %d %s [%s]",
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
			rec.status,
			duration,
			start.Format(time.RFC3339),
		)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}
