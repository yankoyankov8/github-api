package apis

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type ResponseRepoInfo struct {
	CreatedAt   string `json:"created_at"`
	Description string `json:"description"`
	UpdatedAt   string `json:"updated_at"`
	URL         string `json:"url"`
}

type ResponseCommitInfo struct {
	Commit struct {
		Author struct {
			Date  string `json:"date"`
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"author"`
		CommentCount int `json:"comment_count"`
		Committer    struct {
			Date  string `json:"date"`
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"committer"`
		Message string `json:"message"`
		Tree    struct {
			Sha string `json:"sha"`
			URL string `json:"url"`
		} `json:"tree"`
		URL          string `json:"url"`
		Verification struct {
			Payload   interface{} `json:"payload"`
			Reason    string      `json:"reason"`
			Signature interface{} `json:"signature"`
			Verified  bool        `json:"verified"`
		} `json:"verification"`
	} `json:"commit"`
}

type GithubApi struct {
	BaseURL *url.URL
}

type GithubApiInterface interface {
	GetRepoInfo(repo string) (ResponseRepoInfo, error)
	GetCommitInfo(repo string, commit string) (ResponseCommitInfo, error)
}

func (g *GithubApi) GetRepoInfo(repo string) (ResponseRepoInfo, error) {
	if repo == "" {
		return ResponseRepoInfo{}, errors.New("Repo param is required!")
	}

	rel := &url.URL{Path: fmt.Sprintf("repos/vmware/%s", repo)}
	u := g.BaseURL.ResolveReference(rel)

	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		fmt.Errorf(err.Error())
		return ResponseRepoInfo{}, err
	}

	resp, err := client.Do(request)
	defer resp.Body.Close()
	if err != nil {
		fmt.Errorf(err.Error())
		return ResponseRepoInfo{}, err
	}

	var result ResponseRepoInfo
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}

func (g *GithubApi) GetCommitInfo(repo string, commit string) (ResponseCommitInfo, error) {
	if repo == "" || commit == "" {
		return ResponseCommitInfo{}, errors.New("Repo and commit params are required!")
	}

	rel := &url.URL{Path: fmt.Sprintf("repos/vmware/%s/commits/%s", repo, commit)}
	u := g.BaseURL.ResolveReference(rel)

	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		fmt.Errorf(err.Error())
		return ResponseCommitInfo{}, err
	}

	resp, err := client.Do(request)
	defer resp.Body.Close()
	if err != nil {
		fmt.Errorf(err.Error())
		return ResponseCommitInfo{}, err
	}

	var result ResponseCommitInfo
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}
