package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("Access from ", r.URL.Path)

		h.ServeHTTP(w, r)
	})
}
