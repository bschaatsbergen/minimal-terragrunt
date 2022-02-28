inputs = {
  environment = "dev"
}

generate "provider" {
  path      = "provider.tf"
  if_exists = "overwrite_terragrunt"
  contents  = <<EOF
provider "aws" {
  region  = "eu-central-1"
}
EOF
}

remote_state {
  backend = "s3"
  config = {
    bucket         = get_env("BACKEND_S3_BUCKET_DEV")
    key            = "${get_aws_account_id()}/${path_relative_to_include()}/terraform.tfstate"
    region         = get_env("BACKEND_REGION")
    encrypt        = true
    dynamodb_table = get_env("BACKEND_DYNAMODB_TABLE_DEV")
  }
  generate = {
    path      = "backend.tf"
    if_exists = "overwrite_terragrunt"
  }
}
