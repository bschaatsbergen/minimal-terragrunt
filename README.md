# Minimal Terragrunt

Terragrunt is a thin-wrapper around Terraform with a few extra features to keep your code 'dry'.

## What is minimal terragrunt?

Minimal Terragrunt is a project that you can refer to as a 'minimal' implementation of a terragrunt based project.
In this project we make use of the most common features of terragrunt and make it as simple and cross-translatable as possible.

## Modules directory

When you make use of terragrunt it's a good practice that you develop isolated autonomous modules. You simply develop a set of modules that you can pick and choose from and apply to different environments. Our `modules` directory is a good example of this, it contains a `iam` and `s3` directory. Both directories contain a set of AWS resources.


## Environments directory

The `environments` directory is where you define your environments. Each environment is a directory that contains subdirectories for each module.
You can pick whatever module is appropriate for your environment. Every subdirectory contains a `terragrunt.hcl` file that orchestrates the deployment of a particular module.

## Refer a module

Every `terragrunt.hcl` file has a reference to a module through:

```hcl
terraform {
  source = "${get_path_to_repo_root()}/modules/s3"
}
```

The `get_path_to_repo_root()` function is [a simple helper function](https://terragrunt.gruntwork.io/docs/reference/built-in-functions/#get_path_to_repo_root) that returns the path to the root of the repository.


## Inheritting input variables

Terragrunt allows you to inherit input variables from a parent directory (it can only inherit once).
I personally prefer to declare a `terragrunt.hcl` file with input variables in the root directory of a environment, e.g. the `/environments/dev` directory.

```hcl
inputs = {
  environment = "dev"
}
```

The input is injected as values for your module its variables. In both our `iam` and `s3` modules we have a `environment` variable, I therefore declare its value in the `/environments/dev/terragrunt.hcl` file, and inherit it from both the `environments/dev/iam` `environments/dev/s3` directory.

If I have a value for a module that isn't shared with other modules, I can declare it in the `/environments/dev/iam/terragrunt.hcl` file, e.g. the `default_bucket_arn` input I declared.

## Dependencies and passing values to modules

My favorite way to pass values to modules is to use Terragrunt. You can pass values to modules by using the dependency syntax.
In the `/environments/dev/iam/terragrunt.hcl` file we declare a dependency on the `s3` module, and pass its output to the `iam` module.

```hcl
inputs = {
  default_bucket_arn = dependency.s3.outputs.default_bucket_arn
}

dependency "s3" {
  config_path = "${get_path_to_repo_root()}/environments/dev/s3"

  mock_outputs = {
    default_bucket_arn = "arn:aws:s3:::example-bucket"
  }
}
```

The mock_outputs are used to mock the outputs of the `s3` module, so that we can test the `iam` module without having to create a real S3 bucket.


## Contribute

If you want to contribute to this project, please open an issue or pull request.