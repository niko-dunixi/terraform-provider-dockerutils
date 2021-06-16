package main

import (
	"context"

	docker "github.com/docker/docker/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return Provider()
		},
	})
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			// "host": {
			// 	Type:        schema.TypeString,
			// 	Optional:    true,
			// 	DefaultFunc: schema.EnvDefaultFunc("DOCKER_HOST", nil),
			// 	Description: "The TCP or UNIX socket for communication with the Docker daemon",
			// },
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ResourcesMap:         map[string]*schema.Resource{},
		ConfigureContextFunc: configureProvider,
	}
}

func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	dockerClient, err := docker.NewClientWithOpts(docker.FromEnv, docker.WithAPIVersionNegotiation())
	if err != nil {
		return nil, diag.Errorf("could not initialize docker client: %+v", err)
	}
	if _, err := dockerClient.Ping(ctx); err != nil {
		return nil, diag.Errorf("could not ping docker daemon/server: %+v", err)
	}
	return &DockerUtilsConfig{
		dockerClient: dockerClient,
	}, nil
}
