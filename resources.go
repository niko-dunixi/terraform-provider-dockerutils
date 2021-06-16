package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DockerBuild() *schema.Resource {
	return &schema.Resource{}
}

// func DockerImageDigests() *schema.Resource {
// 	return &schema.Resource{
// 		ReadContext: readDockerImageDigests,
// 		Schema: map[string]*schema.Schema{
// 			"name": {
// 				Type:        schema.TypeString,
// 				Description: "The name of the Docker image",
// 				Required:    true,
// 				ForceNew:    true,
// 			},
// 		},
// 		// Schema: map[string]*schema.Schema{
// 		// 	"digests": &schema.Schema{
// 		// 		Type:     schema.TypeList,
// 		// 		Computed: true,
// 		// 		Elem: &schema.Schema{
// 		// 			Type: schema.TypeString,
// 		// 		},
// 		// 	},
// 		// },
// 	}
// }

// func readDockerImageDigests(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	// dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
// 	dockerClient := m.(*DockerUtilsConfig).dockerClient
// 	imageName := d.Get("name").(string)
// 	_, err := dockerClient.ImagePull(ctx, imageName, types.ImagePullOptions{})
// 	if err != nil {
// 		diag.FromErr(err)
// 	}
// 	hist, err := dockerClient.ImageHistory(ctx, imageName)
// 	return nil
// }

// // func newHttpClient() *http.Client {
// 	return &http.Client{
// 		Timeout: 10 * time.Second,
// 	}
// }
