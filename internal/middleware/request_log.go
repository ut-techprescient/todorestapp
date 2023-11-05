package middleware

import (
	"net/http"

	"todorestapp.ut/cmd/config"
)

func RequestLoggerMiddleWare(appConfig *config.AppEnv) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// MiddleWare Logic
			appConfig.InfoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

			next.ServeHTTP(w, r)
		})
	}
}
