package handlers

import (
	"api/apis"
	"api/helpers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

//TODO Not for here!
var users = map[string]string{
	"user1": "password1",
}

func ValidateJwtToken(w http.ResponseWriter, r *http.Request, jwtClaims helpers.JwtInterface) {
	// Get the JSON body and decode into token
	type post struct {
		Token string `json:"token"`
	}
	var postData post
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"error":"Wrong json format."}`)
		return
	}

	if postData.Token == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"error":"missing token"}`)
		return
	}

	// Parse the JWT string and store the result in `claims`.
	err, isValid, username := jwtClaims.IsTokenValid(postData.Token)
	if err != nil || isValid == false {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"error":"Not valid token"}`)
		return
	}

	type result struct {
		Result string
	}
	res := result{Result: fmt.Sprintf("Token is valid for user: %s!", username)}
	js, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":"Internal error"}`)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func IssuingJwtToken(w http.ResponseWriter, r *http.Request, jwtClaims helpers.JwtInterface) {
	// Get the JSON body and decode into credentials
	type post struct {
		Password string `json:"password"`
		Username string `json:"username"`
	}
	var postData post
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"error":"Wrong json format"}`)
		return
	}

	// Get the expected password
	expectedPassword, ok := users[postData.Username]

	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	if !ok || expectedPassword != postData.Password {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"error":"Wrong user/password."}`)
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 2 hours
	expirationTime := time.Now().Add(2 * time.Hour).Unix()
	// Create the JWT claims, which includes the username and expiry time
	//jwtClaims = &helpers.Claims{
	//	Username: credentials.Username,
	//	StandardClaims: jwt.StandardClaims{
	//		// In JWT, the expiry time is expressed as unix milliseconds
	//		ExpiresAt: expirationTime.Unix(),
	//	},
	//}
	err, token := jwtClaims.GenerateToken(postData.Username, expirationTime)

	type result struct {
		Token string
	}
	res := result{Token: token}
	js, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":"Internal error"}`)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetFunctionalities(w http.ResponseWriter, r *http.Request) {

	functionalities := `{
	"functions":{
		"GetIssuingJwtToken":{"url":"/issuing-jwt-token", "request-type":"POST", "accept-post-json":{"username":"user1", "password":"password1"}, "response":"json"},
		"GetValidateJwtToken":{"url":"/validate-jwt-token", "request-type":"POST", "accept-post-json":{"token":""}, "response":"json"}, 
		"GetRepoInformation":{"url":"/get-repo-information", "request-type":"POST", "accept-post-json":{"token":"", "repo":""}, "response":"json"}, 
		"GetCommitInformation":{"url":"/get-commit-information", "request-type":"POST", "accept-post-json":{"token":"", "repo":"", "commit":""}, "response":"json"}
	}
}`
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(functionalities))
}

func GetRepoInformation(w http.ResponseWriter, r *http.Request, github apis.GithubApiInterface) {
	// Get the JSON body and decode
	type post struct {
		Repo string `json:"repo"`
	}
	var postData post
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		fmt.Errorf(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"error":"Decoding Json."}`)
		return
	}
	if postData.Repo == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"error":"Repo is required!"}`)
		return
	}

	response, err := github.GetRepoInfo(postData.Repo)
	if err != nil {
		fmt.Errorf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":"Internal error"}`)
		return
	}

	js, err := json.Marshal(response)
	if err != nil {
		fmt.Errorf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":"Internal error"}`)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetCommitInformation(w http.ResponseWriter, r *http.Request, github apis.GithubApiInterface) {
	// Get the JSON body and decode
	type post struct {
		Repo   string `json:"repo"`
		Commit string `json:"commit"`
	}
	var postData post
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"error":"Wrong json format."}`)
		return
	}
	if postData.Repo == "" || postData.Commit == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"error":"Repo and Commit are required!"}`)
		return
	}

	response, err := github.GetCommitInfo(postData.Repo, postData.Commit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":"Internal error"}`)
		return
	}

	js, err := json.Marshal(response)
	if err != nil {
		fmt.Errorf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":"Internal error"}`)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
