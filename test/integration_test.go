package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	tfOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "./simple",
		EnvVars: map[string]string{
			"TF_LOG_PROVIDER": "TRACE",
		},
	})
	defer terraform.Destroy(t, tfOptions)

	terraform.InitAndApply(t, tfOptions)
	tfOutputHelloWorldStatic := terraform.Output(t, tfOptions, "hello_world_static")
	assert.Equal(t, "STATIC: Hello, World!", tfOutputHelloWorldStatic)
	tfOutputHelloWorld := terraform.Output(t, tfOptions, "hello_world")
	assert.Equal(t, "COMPUTED: Hello, World!", tfOutputHelloWorld)
	terraform.InitAndApply(t, tfOptions)
	tfOutputHelloPhil := terraform.Output(t, tfOptions, "hello_phil")
	assert.Equal(t, "COMPUTED: Hello, Phil!", tfOutputHelloPhil)
}
