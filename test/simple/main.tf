terraform {
  required_version = ">= 0.14"
  required_providers {
    dockerutils = {
      source = "paulfreaknbaker.com/providers/dockerutils"
      version = "0.0.0-testing"
    }
  }
}

provider "dockerutils" {
}

data "dockerutils_helloworld" "world" {
}

data "dockerutils_helloworld" "phil" {
  name = "Phil"
}

output "hello_world_static" {
  value = "STATIC: Hello, World!"
}

output "hello_world" {
  value = "COMPUTED: ${data.dockerutils_helloworld.world.greeting}"
}

output "hello_phil" {
  value = "COMPUTED: ${data.dockerutils_helloworld.phil.greeting}"
}