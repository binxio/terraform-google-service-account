locals {
  project     = var.project
  environment = "all"

  service_accounts = {
    "bucket_reader" = {}
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
