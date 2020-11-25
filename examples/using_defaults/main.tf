locals {
  project     = "demo"
  environment = "dev"

  service_accounts = {
    "bucket_reader" = {}
  }
}

module "sa" {
  source  = "binxio/service-account/google"
  version = "~> 1.0.0"

  project     = local.project
  environment = local.environment

  service_accounts = local.service_accounts
}
