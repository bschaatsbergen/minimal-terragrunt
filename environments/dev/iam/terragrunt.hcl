include {
  path = find_in_parent_folders()
}

terraform {
  source = "${get_path_to_repo_root()}/modules/iam"
}

inputs = {
  default_bucket_arn = dependency.s3.outputs.default_bucket_arn
}

dependency "s3" {
  config_path = "${get_path_to_repo_root()}/environments/dev/s3"

  mock_outputs = {
    default_bucket_arn = "arn:aws:s3:::example-bucket"
  }
}
