package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestIAMModule(t *testing.T) {
	t.Parallel()

	moduleDir := "../environments/dev/iam"

	// Clean up any existing `.terragrunt-cache` directory
	deleteTerragruntCache(t, moduleDir)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformBinary: "terragrunt",
		TerraformDir:    moduleDir,
		BackendConfig: map[string]interface{}{
			"bucket":         BACKEND_TEST_S3_BUCKET_NAME,
			"key":            fmt.Sprintf("%s/clear.tfstate", uniqueID()),
			"region":         BACKEND_TEST_REGION,
			"encrypt":        true,
			"dynamodb_table": BACKEND_TEST_DDB_TABLE_NAME,
		},
		Reconfigure: true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Clean up after ourselves to prevent leftover state in the `.terragrunt-cache` directory
	deleteTerragruntCache(t, moduleDir)
}