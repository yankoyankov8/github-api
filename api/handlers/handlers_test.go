package handlers

import (
	"api/apis"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"errors"
)

//Mock Jwt functions
type mockJwt struct{}

func (m mockJwt) GenerateToken(username string, expireTime int64) (error, string) {
	return nil, "test_token"
}
func (m mockJwt) IsTokenValid(token string) (error, bool, string) {
	return nil, true, "TestUser"
}

//Mock Gitgub functions
type mockGithub struct{}

func (m mockGithub) GetRepoInfo(repo string) (apis.ResponseRepoInfo, error) {
	return apis.ResponseRepoInfo{Description: "test"}, nil
}
func (m mockGithub) GetCommitInfo(repo string, commit string) (apis.ResponseCommitInfo, error) {
	return apis.ResponseCommitInfo{}, nil
}


//Mock Jwt functions return Error
type mockJwtError struct{}

func (m mockJwtError) GenerateToken(username string, expireTime int64) (error, string) {
	return errors.New("test"), ""
}
func (m mockJwtError) IsTokenValid(token string) (error, bool, string) {
	return errors.New("test"), false, ""
}


//Mock Gitgub functions return Error
type mockGithubError struct{}

func (m mockGithubError) GetRepoInfo(repo string) (apis.ResponseRepoInfo, error) {
	return apis.ResponseRepoInfo{Description: "test"}, nil
}
func (m mockGithubError) GetCommitInfo(repo string, commit string) (apis.ResponseCommitInfo, error) {
	return apis.ResponseCommitInfo{}, nil
}

//-------------Tests GetCommitInformation handler---------------------
func TestGetCommitInformation(t *testing.T) {
	req, _ := http.NewRequest("POST", "/get-commit-information", strings.NewReader(`{"token":"test", "repo":"admiral", "commit":"553aa1e036f04c414a84ad234586a35dd5845db1"}`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetCommitInformation(w, r, mockGithub{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "{\"commit\":{\"author\":{\"date\":\"\",\"email\":\"\",\"name\":\"\"},\"comment_count\":0,\"committer\":{\"date\":\"\",\"email\":\"\",\"name\":\"\"},\"message\":\"\",\"tree\":{\"sha\":\"\",\"url\":\"\"},\"url\":\"\",\"verification\":{\"payload\":null,\"reason\":\"\",\"signature\":null,\"verified\":false}}}", rr.Body.String())
}

func TestGetCommitInformationMissingRepo(t *testing.T) {
	req, _ := http.NewRequest("POST", "/get-commit-information", strings.NewReader(`{"token":"test", "commit":"553aa1e036f04c414a84ad234586a35dd5845db1"}`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetCommitInformation(w, r, mockGithub{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "{\"error\":\"Repo and Commit are required!\"}", rr.Body.String())
}
func TestGetCommitInformationMissingCommit(t *testing.T) {
	req, _ := http.NewRequest("POST", "/get-commit-information", strings.NewReader(`{"token":"test", "repo":"admiral"}`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetCommitInformation(w, r, mockGithub{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "{\"error\":\"Repo and Commit are required!\"}", rr.Body.String())
}
func TestGetCommitInformationEmptyCommit(t *testing.T) {
	req, _ := http.NewRequest("POST", "/get-commit-information", strings.NewReader(`{"token":"test", "repo":"admiral", "commit":""}`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetCommitInformation(w, r, mockGithub{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "{\"error\":\"Repo and Commit are required!\"}", rr.Body.String())
}

func TestGetCommitInformationEmptyRepo(t *testing.T) {
	req, _ := http.NewRequest("POST", "/get-commit-information", strings.NewReader(`{"token":"test", "repo":"", "commit":"553aa1e036f04c414a84ad234586a35dd5845db1"}`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetCommitInformation(w, r, mockGithub{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "{\"error\":\"Repo and Commit are required!\"}", rr.Body.String())
}
func TestGetCommitInformationWrongJson(t *testing.T) {
	req, _ := http.NewRequest("POST", "/get-commit-information", strings.NewReader(`{"token":"test", "repo":"", "commit":"553aa1e036f04c414a84ad234586a35dd5845db1"`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetCommitInformation(w, r, mockGithub{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "{\"error\":\"Wrong json format.\"}", rr.Body.String())
}

//-------------Tests GetRepoInformation handler---------------------
func TestGetRepoInformation(t *testing.T) {
	req, _ := http.NewRequest("POST", "/get-repo-information", strings.NewReader(`{"token":"test", "repo":"admiral"}`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetRepoInformation(w, r, mockGithub{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "{\"created_at\":\"\",\"description\":\"test\",\"updated_at\":\"\",\"url\":\"\"}", rr.Body.String())
}

func TestGetRepoInformationEmptyRepo(t *testing.T) {
	req, _ := http.NewRequest("POST", "/get-repo-information", strings.NewReader(`{"token":"test", "repo":""}`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetRepoInformation(w, r, mockGithub{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "{\"error\":\"Repo is required!\"}", rr.Body.String())
}

func TestGetRepoInformationMissingRepo(t *testing.T) {
	req, _ := http.NewRequest("POST", "/get-repo-information", strings.NewReader(`{"token":"test"}`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetRepoInformation(w, r, mockGithub{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "{\"error\":\"Repo is required!\"}", rr.Body.String())
}

func TestGetRepoInformationWrongJson(t *testing.T) {
	req, _ := http.NewRequest("POST", "/get-repo-information", strings.NewReader(`{"token":"test"`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetRepoInformation(w, r, mockGithub{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "{\"error\":\"Decoding Json.\"}", rr.Body.String())
}

//------------------Tests issuing token handler------------------------
func TestIssuingJwtToken(t *testing.T) {
	req, _ := http.NewRequest("POST", "/issuing-jwt-token", strings.NewReader(`{"username":"user1", "password":"password1"}`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		IssuingJwtToken(w, r, mockJwt{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, `{"Token":"test_token"}`, rr.Body.String())
}

func TestIssuingJwtTokenWrongJson(t *testing.T) {
	req, _ := http.NewRequest("POST", "/issuing-jwt-token", strings.NewReader(`{"use, "password":"password1"}`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		IssuingJwtToken(w, r, mockJwt{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, `{"error":"Wrong json format"}`, rr.Body.String())
}
func TestIssuingJwtTokenEmptyPassword(t *testing.T) {
	req, _ := http.NewRequest("POST", "/issuing-jwt-token", strings.NewReader(`{"username":"test", "password":""}`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		IssuingJwtToken(w, r, mockJwt{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Equal(t, "{\"error\":\"Wrong user/password.\"}", rr.Body.String())
}
func TestIssuingJwtTokenMissingPassword(t *testing.T) {
	req, _ := http.NewRequest("POST", "/issuing-jwt-token", strings.NewReader(`{"username":"test"}`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		IssuingJwtToken(w, r, mockJwt{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Equal(t, "{\"error\":\"Wrong user/password.\"}", rr.Body.String())
}
func TestIssuingJwtTokenMissingUsername(t *testing.T) {
	req, _ := http.NewRequest("POST", "/issuing-jwt-token", strings.NewReader(`{"password":"test"}`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		IssuingJwtToken(w, r, mockJwt{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Equal(t, "{\"error\":\"Wrong user/password.\"}", rr.Body.String())
}
func TestIssuingJwtTokenRmptyUsername(t *testing.T) {
	req, _ := http.NewRequest("POST", "/issuing-jwt-token", strings.NewReader(`{"username":""}`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		IssuingJwtToken(w, r, mockJwt{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Equal(t, "{\"error\":\"Wrong user/password.\"}", rr.Body.String())
}

//------------------Tests validate token handler---------------
func TestValidateJwtToken(t *testing.T) {
	req, _ := http.NewRequest("POST", "/validate-jwt-token", strings.NewReader(`{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwiZXhwIjoxNTc1MTE2MTEwfQ.KkqrXkiPFYpdCS4OyCNawTGtgWViZXPM-_GAHyeX7DY"}`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateJwtToken(w, r, mockJwt{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, `{"Result":"Token is valid for user: TestUser!"}`, rr.Body.String())
}
func TestValidateJwtEmptyToken(t *testing.T) {
	req, _ := http.NewRequest("POST", "/validate-jwt-token", strings.NewReader(`{"token": ""}`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateJwtToken(w, r, mockJwt{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, http.StatusBadRequest)
	assert.Equal(t, `{"error":"missing token"}`, rr.Body.String())
}
func TestValidateJwtMissingToken(t *testing.T) {
	req, _ := http.NewRequest("POST", "/validate-jwt-token", strings.NewReader(`{}`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateJwtToken(w, r, mockJwt{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, http.StatusBadRequest)
	assert.Equal(t, `{"error":"missing token"}`, rr.Body.String())
}
func TestValidateJwtWrongJson(t *testing.T) {
	req, _ := http.NewRequest("POST", "/validate-jwt-token", strings.NewReader(`{"token": ""`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateJwtToken(w, r, mockJwt{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, http.StatusBadRequest)
	assert.Equal(t, `{"error":"Wrong json format."}`, rr.Body.String())
}

//-------------------Test GetFunctionalities-----------
func TestGetFunctionalities(t *testing.T) {
	req, _ := http.NewRequest("POST", "/get-functionalities", strings.NewReader(`{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzMDA3ODh9.K1MWq28W3cak4niCP9QLnJB-Qr8vzH6GAA7du7_ofTU"`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetFunctionalities(w, r)
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code)
}


//Mock Jwt and Github functions return errors
func TestIssuingJwtTokenErr(t *testing.T) {
	req, _ := http.NewRequest("POST", "/issuing-jwt-token", strings.NewReader(`{"username":"user1", "password":"password1"}`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		IssuingJwtToken(w, r, mockJwtError{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t,  "{\"Token\":\"\"}", rr.Body.String())
}


//------------------Tests validate token handler---------------
func TestValidateJwtTokenErr(t *testing.T) {
	req, _ := http.NewRequest("POST", "/validate-jwt-token", strings.NewReader(`{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwiZXhwIjoxNTc1MTE2MTEwfQ.KkqrXkiPFYpdCS4OyCNawTGtgWViZXPM-_GAHyeX7DY"}`))

	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateJwtToken(w, r, mockJwtError{})
	})

	testHandler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, http.StatusUnauthorized)
	assert.Equal(t, "{\"error\":\"Not valid token\"}", rr.Body.String())
}