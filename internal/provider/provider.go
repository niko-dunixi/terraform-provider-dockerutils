package provider

import (
	"context"
	"fmt"

	docker "github.com/docker/docker/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			// "host": {
			// 	Type:        schema.TypeString,
			// 	Optional:    true,
			// 	DefaultFunc: schema.EnvDefaultFunc("DOCKER_HOST", nil),
			// 	Description: "The TCP or UNIX socket for communication with the Docker daemon",
			// },
		},
		DataSourcesMap: map[string]*schema.Resource{
			"dockerutils_helloworld": DataSourceHelloWorld(),
		},
		ResourcesMap: map[string]*schema.Resource{},
		// ConfigureFunc: func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		// 	thing := struct{}{}
		// 	return &thing, []diag.Diagnostic{{
		// 		Severity: diag.Warning,
		// 		Detail:   fmt.Sprintf("Configuring provider: %+v", &d),
		// 	}}
		// },
		// ConfigureContextFunc: configureProvider,
		// ConfigureFunc: func(rd *schema.ResourceData) (interface{}, error) {
		// 	providerConfig := &structs.ProviderConfig{}
		// 	diagnostics := diag.Diagnostics{{
		// 		Severity: diag.Warning,
		// 		Summary:  fmt.Sprintf("Configuring Provider: %+v", *rd),
		// 	}, {
		// 		Severity: diag.Warning,
		// 		Summary:  fmt.Sprintf("%+v", rd.Get("key").(string)),
		// 	}}
		// 	return providerConfig, nil
		// },
		ConfigureContextFunc: func(ctx context.Context, rd *schema.ResourceData) (interface{}, diag.Diagnostics) {
			diagnostics := diag.Diagnostics{{
				Severity: diag.Warning,
				Summary:  fmt.Sprintf("Configuring Provider: %+v", *rd),
			}}
			dockerClient, err := docker.NewClientWithOpts(docker.FromEnv, docker.WithAPIVersionNegotiation())
			if err != nil {
				diagnostics = append(diagnostics,
					diag.Errorf("could not initialize docker client: %+v", err)...,
				)
				return nil, diagnostics
			}
			if _, err := dockerClient.Ping(ctx); err != nil {
				diagnostics = append(diagnostics,
					diag.Errorf("could not ping docker daemon/server: %+v", err)...,
				)
				return nil, diagnostics
			}
			providerConfig := ProviderConfig{
				DockerClient: dockerClient,
			}
			return &providerConfig, diagnostics
		},
		// ConfigureContextFunc: func(ctx context.Context, d *schema.ResourceData) (providerConfig interface{}, diagnostics diag.Diagnostics) {
		// 	providerConfig = &structs.ProviderConfig{}
		// 	diagnostics = diag.Diagnostics{{
		// 		Severity: diag.Warning,
		// 		Summary:  fmt.Sprintf("Configuring Provider: %+v", *d),
		// 	}, {
		// 		Severity: diag.Warning,
		// 		Summary:  fmt.Sprintf("%+v", d.Get("key").(string)),
		// 	}}
		// 	return
		// },
	}
}
