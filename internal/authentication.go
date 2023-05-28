package internal

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
	"github.com/maayarosama/Blogging_system/models"
)

type UserIDKey string

func Authentication(excludedRoutes []*mux.Route, secret string, timeout int) func(http.Handler) http.Handler {
	// Cache the regex object of each route (obviously for performance purposes)
	var excludedRoutesRegexp []*regexp.Regexp
	rl := len(excludedRoutes)

	for i := 0; i < rl; i++ {
		r := excludedRoutes[i]
		pathRegexp, _ := r.GetPathRegexp()

		regx, _ := regexp.Compile(pathRegexp)
		excludedRoutesRegexp = append(excludedRoutesRegexp, regx)
	}
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			exclude := false
			requestMethod := r.Method

			for i := 0; i < rl; i++ {
				excludedRoute := excludedRoutes[i]
				methods, _ := excludedRoute.GetMethods()

				ml := len(methods)

				methodMatched := false
				if ml < 1 {
					methodMatched = true
				} else {
					for j := 0; j < ml; j++ {
						if methods[j] == requestMethod {
							methodMatched = true
							break
						}
					}
				}
				if methodMatched {
					uri := r.RequestURI
					if excludedRoutesRegexp[i].MatchString(uri) {
						exclude = true
						break
					}
				}
			}
			if !exclude {
				reqToken := r.Header.Get("Authorization")
				splitToken := strings.Split(reqToken, "Bearer ")

				if len(splitToken) != 2 {

					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("User is not authorized \n"))
					return
				}
				reqToken = splitToken[1]

				claims, err := models.ValidateToken(reqToken, secret, timeout)
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("User is not authorized \n"))
					return
				}
				ctx := context.WithValue(r.Context(), UserIDKey("UserID"), claims.UserID)

				h.ServeHTTP(w, r.WithContext(ctx))
			} else {
				h.ServeHTTP(w, r)
			}
		})
	}
}
