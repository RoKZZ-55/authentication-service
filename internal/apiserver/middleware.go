package apiserver

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) LogRequest() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := slog.With(
				slog.String("ReqID", middleware.GetReqID(r.Context())),
				slog.String("Method", r.Method),
				slog.String("Proto", r.Proto),
				slog.String("URL", r.Host+r.URL.Path),
				slog.String("RemoteAddr", r.RemoteAddr),
				slog.String("UserAgent", r.UserAgent()),
			)

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			start := time.Now()
			defer func() {
				entry.Info(
					"request result",
					slog.Int("status", ww.Status()),
					slog.Int("bytes", ww.BytesWritten()),
					slog.String("duration", time.Since(start).String()),
				)
			}()

			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}
