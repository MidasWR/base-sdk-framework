package midas

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

type MidConfig struct {
	Log           zerolog.Logger
	TokenInHeader bool
	VerifyJWT     func(token string) error
	HeaderIn      bool
}

func Middleware(next http.HandlerFunc, cfg MidConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if cfg.HeaderIn {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Token")
		}
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		start := time.Now()
		if cfg.TokenInHeader {
			token := r.Header.Get("Token")
			if err := cfg.VerifyJWT(token); err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(err.Error()))
				return
			}
		}

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		cfg.Log.Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Str("remote", r.RemoteAddr).
			Dur("duration", duration).
			Msg("HTTP request")
	}
}
