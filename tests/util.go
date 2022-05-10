package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
)

var (
	BACKEND_TEST_S3_BUCKET_NAME = os.Getenv("BACKEND_TEST_S3_BUCKET_NAME")
	BACKEND_TEST_DDB_TABLE_NAME = os.Getenv("BACKEND_TEST_DDB_TABLE_NAME")
	BACKEND_TEST_REGION         = os.Getenv("BACKEND_TEST_REGION")
)

func uniqueID() string {
	return uuid.New().String()[:7]
}

func deleteTerragruntCache(t *testing.T, dir string) {
	t.Log(t, "Deleting terragrunt cache...")
	os.RemoveAll(fmt.Sprintf("%s/.terragrunt-cache", dir))
}
