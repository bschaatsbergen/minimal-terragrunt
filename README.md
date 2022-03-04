# Minimal Terragrunt

Terragrunt is a thin-wrapper around Terraform with a few extra features to keep your code 'dry'.

## What is minimal terragrunt?

Minimal Terragrunt is a project that you can refer to as a 'minimal' implementation of a terragrunt based project.
In this project we make use of the most common features of terragrunt and make it as simple and cross-translatable as possible.

## Modules directory

When you make use of terragrunt it's a good practice that you develop isolated autonomous modules. You simply develop a set of modules that you can pick and choose from and apply to different environments. Our `modules` directory is a good example of this, it contains a `iam` and `s3` directory.


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


