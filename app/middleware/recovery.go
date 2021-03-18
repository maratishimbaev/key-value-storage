package middleware

import (
	"github.com/kataras/golog"
	"net/http"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				golog.Errorf("panic: %s", err.(error).Error())
			}

			w.WriteHeader(http.StatusInternalServerError)
		}()

		next.ServeHTTP(w, r)
	})
}
