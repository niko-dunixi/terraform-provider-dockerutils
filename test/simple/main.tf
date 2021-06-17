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

output "hello_world" {
  value = "Hello, World!"
}