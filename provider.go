package main

import (
	docker "github.com/docker/docker/client"
)

type DockerUtilsConfig struct {
	dockerClient *docker.Client
}

// func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
// 	return
// }
