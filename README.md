
# Module `terraform-google-service-account`

Core Version Constraints:
* `>= 0.13`

Provider Requirements:
* **google (`hashicorp/google`):** (any version)

## Input Variables
* `environment` (required): Company environment for which the resources are created (e.g. dev, tst, acc, prd, all).
* `gcp_project` (required): GCP Project ID override - this is normally not needed and should only be used in tf-projects.
* `project` (required): Company project name.
* `service_accounts` (required): Map of accounts to be created. The key will be used for the SA name so it should describe the SA purpose. A list of roles can be provided to set an authoritive iam policy for the service account.

## Output Values
* `map`: outputs for all service accounts created

## Managed Resources
* `google_service_account.map` from `google`
* `google_service_account_iam_policy.map` from `google`

## Data Resources
* `data.google_iam_policy.map` from `google`

## Creating a new release
After adding your changed and committing the code to GIT, you will need to add a new tag.
```
git tag vx.x.x
git push --tag
```
If your changes might be breaking current implementations of this module, make sure to bump the major version up by 1.

If you want to see which tags are already there, you can use the following command:
```
git tag --list
```
Required APIs
=============
For the VPC services to deploy, the following APIs should be enabled in your project:
 * iam.googleapis.com

Testing
=======
This module comes with [terratest](https://github.com/gruntwork-io/terratest) scripts for both unit testing and integration testing.
A Makefile is provided to run the tests using docker, but you can also run the tests directly on your machine if you have terratest installed.

### Run with make
Make sure to set GOOGLE_CLOUD_PROJECT to the right project and GOOGLE_CREDENTIALS to the right credentials json file
You can now run the tests with docker:
```
make test
```

### Run locally
From the module directory, run:
```
cd test && TF_VAR_owner=$(id -nu) go test
```
