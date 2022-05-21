package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestIAMModule(t *testing.T) {
	moduleDir := "../environments/dev/iam"

	// Clean up any existing `.terragrunt-cache` directory
	defer deleteTerragruntCache(t, moduleDir)

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

	terraform.InitAndApply(t, terraformOptions)

	// Clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)
}
