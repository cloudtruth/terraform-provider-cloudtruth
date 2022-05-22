package cloudtruth

import (
	"context"
	"encoding/json"
	"errors"
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
				Description: "The name of the parameter",
				Type:        schema.TypeString,
				Required:    true,
			},
			"value": &schema.Schema{
				Description: "The parameter's value",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataCloudTruthParameterRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "dataCloudTruthParameterRead")

	// todo: build URL here using using net/url
	env := url.QueryEscape(d.Get("env").(string))
	paramName := url.QueryEscape(d.Get("name").(string))
	project := url.QueryEscape(d.Get("project").(string))
	// todo: determine which query strings to hard code
	paramURL := fmt.Sprintf("%s/projects/%s/parameters/?evaluate=true&mask_secrets=false&name=%s&environment=%s",
		c.config.BaseURL, project, paramName, env)
	tflog.Debug(ctx, paramURL)

	resp, err := c.Get(paramURL)
	fmt.Println(err.Error())
	if err != nil {
		return diag.FromErr(err)
	}

	var params Parameters
	err = json.NewDecoder(resp.Body).Decode(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	// todo:
	// support for a parameters list type
	if params.Count != 1 {
		return diag.FromErr(errors.New(fmt.Sprintf("Param %s not found", paramName)))
	}

	paramValues := params.Results[0].Values
	// The key here is the URL to the specific env/parameter result
	for _, e := range paramValues {
		if e.EnvironmentName == env {
			d.Set("value", e.Value)
			d.SetId(e.ID)
			break // todo: handle more than one value in this map???
		}
	}
	return nil
}

// todo: clean up these structs, we don't need all of the fields
type Parameters struct {
	Count    int       `json:"count"`
	Next     any       `json:"next"`
	Previous any       `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	URL                  string           `json:"url"`
	ID                   string           `json:"id"`
	Name                 string           `json:"name"`
	Description          string           `json:"description"`
	Secret               bool             `json:"secret"`
	Type                 string           `json:"type"`
	Rules                []any            `json:"rules"`
	Project              string           `json:"project"`
	ProjectName          string           `json:"project_name"`
	ReferencingTemplates []any            `json:"referencing_templates"`
	ReferencingValues    []any            `json:"referencing_values"`
	Values               map[string]Value `json:"values"`
	Overrides            any              `json:"overrides"`
	CreatedAt            time.Time        `json:"created_at"`
	ModifiedAt           time.Time        `json:"modified_at"`
	Templates            []any            `json:"templates"`
}

type Value struct {
	URL                  string    `json:"url"`
	ID                   string    `json:"id"`
	Environment          string    `json:"environment"`
	EnvironmentName      string    `json:"environment_name"`
	EarliestTag          any       `json:"earliest_tag"`
	Parameter            string    `json:"parameter"`
	External             bool      `json:"external"`
	ExternalFqn          string    `json:"external_fqn"`
	ExternalFilter       string    `json:"external_filter"`
	ExternalError        any       `json:"external_error"`
	ExternalStatus       any       `json:"external_status"`
	InternalValue        any       `json:"internal_value"`
	Interpolated         bool      `json:"interpolated"`
	Value                string    `json:"value"`
	Evaluated            bool      `json:"evaluated"`
	Secret               bool      `json:"secret"`
	ReferencedParameters []any     `json:"referenced_parameters"`
	ReferencedTemplates  []any     `json:"referenced_templates"`
	CreatedAt            time.Time `json:"created_at"`
	ModifiedAt           time.Time `json:"modified_at"`
	Dynamic              bool      `json:"dynamic"`
	DynamicError         any       `json:"dynamic_error"`
	DynamicFilter        string    `json:"dynamic_filter"`
	DynamicFqn           string    `json:"dynamic_fqn"`
	StaticValue          any       `json:"static_value"`
}
