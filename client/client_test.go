package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRequest_GetBase(t *testing.T) {
	tests := []struct {
		name    string
		baseUrl string
		path    string
		want    string
	}{
		{
			name:    "Test base url ends with /, path starts with /",
			baseUrl: "https://api.example.com/",
			path:    "/v1/account/0x32Be343B94f860124dC4fEe278FDCBD38C102D88",
			want:    "https://api.example.com/v1/account/0x32Be343B94f860124dC4fEe278FDCBD38C102D88",
		},
		{
			name:    "Test only base url ends with /",
			baseUrl: "https://api.example.com/",
			path:    "v1/account/0x32Be343B94f860124dC4fEe278FDCBD38C102D88",
			want:    "https://api.example.com/v1/account/0x32Be343B94f860124dC4fEe278FDCBD38C102D88",
		},
		{
			name:    "Test only path starts with /",
			baseUrl: "https://api.example.com",
			path:    "/v1/account/0x32Be343B94f860124dC4fEe278FDCBD38C102D88",
			want:    "https://api.example.com/v1/account/0x32Be343B94f860124dC4fEe278FDCBD38C102D88",
		},
		{
			name:    "Test none /",
			baseUrl: "https://api.example.com",
			path:    "v1/account/0x32Be343B94f860124dC4fEe278FDCBD38C102D88",
			want:    "https://api.example.com/v1/account/0x32Be343B94f860124dC4fEe278FDCBD38C102D88",
		},
		{
			name:    "Test empty path",
			baseUrl: "https://api.example.com/",
			path:    "",
			want:    "https://api.example.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := InitClient(tt.baseUrl, nil)
			if got := r.GetBase(tt.path); got != tt.want {
				t.Errorf("Request.GetBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_GetURL(t *testing.T) {
	tests := []struct {
		name    string
		baseUrl string
		path    string
		query   url.Values
		want    string
	}{
		{
			name:    "Test empty query",
			baseUrl: "https://3rdparty-apis.coinmarketcap.com",
			path:    "/v1/cryptocurrency/widget?id=1027&convert=USD",
			query:   nil,
			want:    "https://3rdparty-apis.coinmarketcap.com/v1/cryptocurrency/widget?id=1027&convert=USD",
		},
		{
			name:    "Test query",
			baseUrl: "https://3rdparty-apis.coinmarketcap.com",
			path:    "/v1/cryptocurrency/widget",
			query: url.Values{
				"id":      {"1027"},
				"convert": {"USD"},
			},
			want: "https://3rdparty-apis.coinmarketcap.com/v1/cryptocurrency/widget?convert=USD&id=1027",
		},
		{
			name:    "Test query2",
			baseUrl: "https://data.ripple.com/v2",
			path:    "ledgers/61330266",
			query: url.Values{
				"transactions": {"true"},
				"binary":       {"false"},
				"expand":       {"true"},
				"limit":        {"100"},
			},
			want: "https://data.ripple.com/v2/ledgers/61330266?binary=false&expand=true&limit=100&transactions=true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := InitClient(tt.baseUrl, nil)
			if got := r.GetURL(tt.path, tt.query); got != tt.want {
				t.Errorf("Request.GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeoutOption(t *testing.T) {
	tests := []struct {
		name             string
		serverTimeout    time.Duration
		serverResponse   string
		clientTimeout    time.Duration
		expectedResponse string
		errExpected      assert.ErrorAssertionFunc
	}{
		{
			name:           "client exits with timeout err",
			serverTimeout:  time.Millisecond * 6,
			serverResponse: "ok",
			clientTimeout:  time.Millisecond * 2,
			errExpected:    assert.Error,
		},
		{
			name:             "response returned in time",
			serverTimeout:    time.Millisecond * 2,
			serverResponse:   "{\"status\":\"ok\"}",
			clientTimeout:    time.Millisecond * 6,
			expectedResponse: "{\"status\":\"ok\"}",
			errExpected:      assert.NoError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			serverTimeout := tc.serverTimeout
			serverResponse := tc.serverResponse

			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(serverTimeout)

				_, err := fmt.Fprintf(w, serverResponse)
				assert.NoError(t, err)
			}))

			client := InitClient(srv.URL, nil, TimeoutOption(tc.clientTimeout))

			var actual json.RawMessage
			err := client.Get(&actual, "", nil)
			tc.errExpected(t, err)
			assert.Equal(t, tc.expectedResponse, string(actual))
		})
	}
}

type jsonModel struct {
	Name string `json:"name"`
}

type mockJSONClient struct {
	body []byte
}

func newMockJSONClient(b []byte) *mockJSONClient {
	return &mockJSONClient{body: b}
}

func (c *mockJSONClient) Do(_ *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(c.body)),
	}, nil
}
