locals {
  project     = "demo"
  environment = "dev"

  service_accounts = {
    "bucket_reader" = {
      roles = {
        "roles/iam.serviceAccountUser" = {
          (var.sa_user_email) = "serviceAccount"
        }
        "roles/iam.serviceAccountTokenCreator" = {
          (var.sa_user_email) = "serviceAccount"
        }
      }
    }
  }
}

module "sa" {
  source  = "binxio/service-account/google"
  version = "~> 1.0.0"

  project     = local.project
  environment = local.environment

  service_accounts = local.service_accounts
}
