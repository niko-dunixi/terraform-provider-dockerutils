package main

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DockerImageDigests() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDockerImageDigests,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the Docker image",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func readDockerImageDigests(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	dockerClient := m.(*DockerUtilsConfig).dockerClient
	imageName := d.Get("name").(string)
	_, err := dockerClient.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		return diag.FromErr(err)
	}
	response, _, err := dockerClient.ImageInspectWithRaw(ctx, imageName)
	if err != nil {
		return diag.FromErr(err)
	}
	if len(response.RepoDigests) == 0 {
		return diag.Errorf("there were no digests found")
	}
	d.Set("digest", response.RepoDigests[0])
	return nil
}
