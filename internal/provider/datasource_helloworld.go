package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceHelloWorld() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) (diagnostics diag.Diagnostics) {
			inputName := d.Get("name").(string)
			outputGreeting := fmt.Sprintf("Hello, %s!", inputName)
			diagnostics = append(diagnostics, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Input: " + inputName,
			}, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Output: " + outputGreeting,
			})
			if err := (*d).Set("greeting", outputGreeting); err != nil {
				diagnostics = append(diagnostics, diag.FromErr(err)...)
			}
			if err := d.Set("key", "value"); err != nil {
				diagnostics = append(diagnostics, diag.FromErr(err)...)
			}
			d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
			return
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Who to greet",
				Default:     "World",
			},
			"greeting": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
