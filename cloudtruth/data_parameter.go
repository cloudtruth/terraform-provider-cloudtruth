package cloudtruth

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataCloudTruthParameter() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth parameter data source",
		ReadContext: dataCloudTruthParameterRead,
		Schema: map[string]*schema.Schema{
			"env": {
				Description: "The CloudTruth environment",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project": {
				Description: "The CloudTruth project",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "The name of the secret",
				Type:        schema.TypeString,
				Required:    true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//<BaseURL>/projects/<project>/parameters/?evaluate=true&mask_secrets=false&name=<name>&wrap=false&environment=<environment>'
func dataCloudTruthParameterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "dataCloudTruthParameterRead")
	// todo: build URL here using using net/url
	env := url.QueryEscape(d.Get("env").(string))
	name := url.QueryEscape(d.Get("name").(string))
	project := url.QueryEscape(d.Get("project").(string))
	paramURL := fmt.Sprintf("%s/projects/%s/parameters/?evaluate=true&mask_secrets=false&name=%s&environment=%s",
		c.config.BaseURL, project, name, env)
	tflog.Debug(ctx, paramURL)
	resp, err := c.Get(paramURL)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("%v+", resp))
	d.SetId("temporary id")
	return nil
}

type Parameters struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Results   `json:"results"`
}

type Results struct {
	URL                  string           `json:"url"`
	ID                   string           `json:"id"`
	Name                 string           `json:"name"`
	Description          string           `json:"description"`
	Secret               bool             `json:"secret"`
	Type                 string           `json:"type"`
	Rules                []interface{}    `json:"rules"`
	Project              string           `json:"project"`
	ProjectName          string           `json:"project_name"`
	ReferencingTemplates []interface{}    `json:"referencing_templates"`
	ReferencingValues    []interface{}    `json:"referencing_values"`
	Values               map[string]Value `json:"values"`
	Overrides            interface{}      `json:"overrides"`
	CreatedAt            time.Time        `json:"created_at"`
	ModifiedAt           time.Time        `json:"modified_at"`
	Templates            []interface{}    `json:"templates"`
}

type Value struct {
	URL                  string        `json:"url"`
	ID                   string        `json:"id"`
	Environment          string        `json:"environment"`
	EnvironmentName      string        `json:"environment_name"`
	EarliestTag          interface{}   `json:"earliest_tag"`
	Parameter            string        `json:"parameter"`
	External             bool          `json:"external"`
	ExternalFqn          string        `json:"external_fqn"`
	ExternalFilter       string        `json:"external_filter"`
	ExternalError        interface{}   `json:"external_error"`
	ExternalStatus       interface{}   `json:"external_status"`
	InternalValue        interface{}   `json:"internal_value"`
	Interpolated         bool          `json:"interpolated"`
	Value                string        `json:"value"`
	Evaluated            bool          `json:"evaluated"`
	Secret               bool          `json:"secret"`
	ReferencedParameters []interface{} `json:"referenced_parameters"`
	ReferencedTemplates  []interface{} `json:"referenced_templates"`
	CreatedAt            time.Time     `json:"created_at"`
	ModifiedAt           time.Time     `json:"modified_at"`
	Dynamic              bool          `json:"dynamic"`
	DynamicError         interface{}   `json:"dynamic_error"`
	DynamicFilter        string        `json:"dynamic_filter"`
	DynamicFqn           string        `json:"dynamic_fqn"`
	StaticValue          interface{}   `json:"static_value"`
}

