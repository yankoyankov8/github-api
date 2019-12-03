package apis

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
	"fmt"
	"api/helpers"
)

var github = &GithubApi{BaseURL: &url.URL{Host: helpers.GithubUrl, Scheme: "https"}}

func TestGet(t *testing.T) {
	assert := assert.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		description   string
		responder     httpmock.Responder
		expectedRepos ResponseRepoInfo
		expectedError error
	}{
		{
			description: "github api success",
			responder: httpmock.NewStringResponder(200, fmt.Sprintf(`{
    "created_at": "2016-08-22T15:49:01Z",
    "description": "test",
    "updated_at": "2019-11-28T23:58:25Z",
    "url": "https://%s/repos/vmware/admiral"
}`, helpers.GithubUrl)),
			expectedRepos: ResponseRepoInfo{CreatedAt: "2016-08-22T15:49:01Z", Description: "test", UpdatedAt: "2019-11-28T23:58:25Z", URL: "https://api.github.com/repos/vmware/admiral"},
			expectedError: nil,
		},
		{
			description:   "github api success, no repos",
			responder:     httpmock.NewStringResponder(200, `[]`),
			expectedRepos: ResponseRepoInfo{},
			expectedError: nil,
		}, {
			description:   "github api failure, not found",
			responder:     httpmock.NewStringResponder(404, `{"message": "not found"}`),
			expectedRepos: ResponseRepoInfo{},
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		httpmock.RegisterResponder("GET", "https://api.github.com/repos/vmware/test", tc.responder)

		r, err := github.GetRepoInfo("test")

		assert.Equal(r, tc.expectedRepos, tc.description)
		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}
