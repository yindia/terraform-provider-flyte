package flyte

import (
	"context"

	"github.com/flyteorg/flyte/flyteidl/clients/go/admin"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Provider returns the Flyte Terraform Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_ID", nil),
				Description: "FLYTE CLIENT ID",
				Sensitive:   true,
			},
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FLYTE_URL", "https://flyte.org"),
				Description: "Flyte URL (HTTPS) address",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"flyte_project": resourceProject(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	clientSet, err := admin.ClientSetBuilder().WithConfig(&admin.Config{
		Endpoint:           d.Get("url").(string),
		ClientID:           d.Get("client_id").(string),
		ClientSecretEnvVar: "CLIENT_SECRET",
	}).Build(context.Background())
	if err != nil {
		return nil, err
	}
	return clientSet.AdminClient(), nil
}
