package provider

import (
	docker "github.com/docker/docker/client"
)

type ProviderConfig struct {
	DockerClient *docker.Client
}
