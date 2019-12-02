package middlewares

import (
	"api/helpers"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read the POST JSOM content
		type post struct {
			Token string `json:"token"`
		}
		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(r.Body)
		}
		// Restore the io.ReadCloser to its original state
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		postData := post{}
		err := json.Unmarshal(bodyBytes, &postData)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"error":"Json in wrong format."}`)
			return
		}
		if postData.Token == "" {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"error":"missing token"}`)
			return
		}

		// Parse the JWT string and store the result in `claims`.
		claims := &helpers.Claims{}
		err, isValid, _ := claims.IsTokenValid(postData.Token)

		if err != nil || isValid != true {
			// Report Unauthorized
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, `{"error":"Invalid token."}`)
			return
		}

		next.ServeHTTP(w, r)
	})
}
