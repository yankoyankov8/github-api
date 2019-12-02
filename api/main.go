package main

import (
	"api/apis"
	"api/handlers"
	"api/helpers"
	"api/middlewares"
	"net/http"
	"net/url"
)

func main() {
	//init Jwt dependency
	jwtClaims := &helpers.Claims{}
	//init Github Api dependency
	githubApi := &apis.GithubApi{BaseURL: &url.URL{Host: "api.github.com", Scheme: "https"}}

	//Init listeners
	http.Handle("/get-functionalities", Middleware(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				handlers.GetFunctionalities(w, r)
			}),
		middlewares.AuthMiddleware,
	))
	http.Handle("/get-repo-information", Middleware(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				handlers.GetRepoInformation(w, r, githubApi)
			}),
		middlewares.AuthMiddleware,
	))
	http.Handle("/get-commit-information", Middleware(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				handlers.GetCommitInformation(w, r, githubApi)
			}),
		middlewares.AuthMiddleware,
	))

	http.HandleFunc("/validate-jwt-token", func(w http.ResponseWriter, r *http.Request) {
		handlers.ValidateJwtToken(w, r, jwtClaims)
	})
	http.HandleFunc("/issuing-jwt-token", func(w http.ResponseWriter, r *http.Request) {
		handlers.IssuingJwtToken(w, r, jwtClaims)
	})
	http.ListenAndServe(":8080", nil)
}

// Middleware (this function) makes adding more than one layer of middleware easy
// by specifying them as a list. It will run the last specified handler first.
func Middleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}
