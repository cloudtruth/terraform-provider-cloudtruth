package cloudtruth

import (
	"fmt"
	"io"
	"net/http"
	//"github.com/hashicorp/terraform-plugin-log/tflog"
	//"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type cloudTruthClient struct {
	client http.Client
	config clientConfig
}

func (c *cloudTruthClient) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err

	}
	return c.Do(req)
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
