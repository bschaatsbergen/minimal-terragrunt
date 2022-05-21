package test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestS3Module(t *testing.T) {
	moduleDir := "../environments/dev/s3"

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
	bucketID := terraform.Output(t, terraformOptions, "bucket_id")

	// Clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Create an S3 session
	s3client := s3.New(session.New(
		&aws.Config{
			Region: aws.String("eu-central-1"),
		},
	))

	// Check that the bucket exists and versioning is enabled
	bucketVersioning, err := s3client.GetBucketVersioning(&s3.GetBucketVersioningInput{
		Bucket: aws.String(bucketID),
	})
	if err != nil {
		t.Fatal(err)
	}
	if *bucketVersioning.Status != "Enabled" {
		t.Fatal("Expected bucket versioning to be enabled")
	}

	// Check that the bucket is encrypted with AES256
	bucketEncryption, err := s3client.GetBucketEncryption(&s3.GetBucketEncryptionInput{
		Bucket: aws.String(bucketID),
	})
	if err != nil {
		t.Fatal(err)
	}
	if *bucketEncryption.ServerSideEncryptionConfiguration.Rules[0].ApplyServerSideEncryptionByDefault.SSEAlgorithm != "AES256" {
		t.Fatal("Expected bucket encryption to be AES256")
	}
}
