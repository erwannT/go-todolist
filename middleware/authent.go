package middleware

import "net/http"

func AuthenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, userNameOk := r.URL.Query()["username"]
		password, passwordOk := r.URL.Query()["password"]

		if userNameOk && passwordOk {

			if username[0] == "erwann" && password[0] == "password" {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "not authenticatd", 403)
			}
		} else {
			http.Error(w, "not authenticatd", 401)
		}
	})
}
