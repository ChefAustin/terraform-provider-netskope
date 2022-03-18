package netskope

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ns-sbrown/nsgo"
)

func dataSourcePublishersRead(d *schema.ResourceData, m interface{}) error {
	//Init a client instance
	nsclient := m.(*nsgo.Client)

	//Get Publisher
	pubs, err := nsclient.GetPublishers()
	if err != nil {
		return err
	}

	if err := d.Set("publishers", pubs.Publishers); err != nil {
		return err
	}

	return nil

}

func dataSourcePublishers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePublishersRead,
		Schema: map[string]*schema.Schema{
			"publishers": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"common_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"registered": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"stitcher_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"assessment": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"eee_support": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"hdd_free": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"hdd_total": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_address": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"version": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
