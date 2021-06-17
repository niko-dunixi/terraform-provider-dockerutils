package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	tfOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "./simple",
	})
	defer terraform.Destroy(t, tfOptions)

	terraform.InitAndApply(t, tfOptions)
	tfOutput := terraform.Output(t, tfOptions, "hello_world")
	assert.Equal(t, "Hello, World!", tfOutput)
}
