package localfile

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
		},
		ResourcesMap: map[string]*schema.Resource{
			"localfile_file": resourceLocalFile(),
		},
		ConfigureFunc: generateConfig,
	}
}

func generateConfig(rd *schema.ResourceData) (interface{}, error) {
	return nil, nil
}
