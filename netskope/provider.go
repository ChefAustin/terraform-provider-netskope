package netskope

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ns-sbrown/nsgo"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"baseurl": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NS_BaseURL", nil),
			},
			"apitoken": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("NS_ApiToken", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"netskope_publisher":   resourcePublishers(),
			"netskope_privateapps": resourcePrivateApps(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"netskope_publisher": dataSourcePublishers(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	nsclient := nsgo.NewClient(d.Get("baseurl").(string), d.Get("apitoken").(string))
	return nsclient, nil
}
