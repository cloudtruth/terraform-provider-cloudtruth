package cloudtruth

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/ctapi"
)

type cloudTruthClient struct {
	client http.Client
	config clientConfig
}

func (c *cloudTruthClient) Get(url string) (resp *http.Response, err error) {
	cfg := ctapi.NewConfiguration()
	fmt.Printf("%v+", cfg)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err

	}
	resp, respErr := c.Do(req)
	if err != nil {
		return nil, respErr
	}
	respErr = checkStatusCodeForError(resp.StatusCode)
	return resp, respErr
}

func (c *cloudTruthClient) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.Do(req)
}

func (c *cloudTruthClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("Api-Key %s", c.config.APIKey))
	req.Header.Add("Accept", "application/json")
	return c.client.Do(req)
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s [HTTP code %d]", e.Type, e.Message, e.StatusCode)
}

// Error represents an error returned by the CloudTruth API.
type Error struct {
	Type       string `json:"type"`
	Message    string `json:"message"`
	StatusCode int
}

// Useful HTTP errors
func checkStatusCodeForError(statusCode int) error {
	if statusCode == 200 {
		return nil
	}
	switch statusCode {
	case 400:
		return Error{
			Type:       "BAD_REQUEST",
			Message:    "The HTTP request is invalid",
			StatusCode: statusCode,
		}
	case 401:
		return Error{
			Type:       "AUTHENTICATION_REQUIRED",
			Message:    "You must provide a valid api key to perform this operation",
			StatusCode: statusCode,
		}
	case 403:
		return Error{
			Type:       "NOT_AUTHORIZED",
			Message:    "You are not authorized to perform this operation",
			StatusCode: statusCode,
		}
	case 404:
		return Error{
			Type:       "NOT_FOUND",
			Message:    "Could not find what you are looking for",
			StatusCode: statusCode,
		}
	case 413:
		return Error{
			Type:       "REQUEST_TOO_LARGE",
			Message:    "The request body is too large",
			StatusCode: statusCode,
		}
	case 422:
		return Error{
			Type:       "UNPROCESSABLE_ENTITY",
			Message:    "CloudTruth was unable to process this request",
			StatusCode: statusCode,
		}
	case 429:
		return Error{
			Type:       "TOO_MANY_REQUESTS",
			Message:    "CloudTruth rate limiting error",
			StatusCode: statusCode,
		}
	case 500:
		return Error{
			Type:       "SERVER_ERROR",
			Message:    "Try again. If the problem persists, contact CloudTruth support.",
			StatusCode: statusCode,
		}
	case 503:
		return Error{
			Type:       "SERVICE_UNAVAILABLE",
			Message:    "CloudTruth is temporarily unavailable. Please retry shortly.",
			StatusCode: statusCode,
		}
	default:
		return Error{
			Type:       "UNKNOWN_HTTP_ERROR",
			Message:    "CloudTruth returned an unknown HTTP response.",
			StatusCode: statusCode,
		}
	}
}

// For debugging purposes, this function generates ascii representation of a request
func formatRequest(r *http.Request) string {
	var request []string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	return strings.Join(request, "\n")
}
