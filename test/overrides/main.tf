locals {
  project     = var.project
  environment = "all"

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
  source = "../../"

  project     = local.project
  environment = local.environment

  service_accounts = local.service_accounts
}

output "map" {
  value = module.sa.map
}
