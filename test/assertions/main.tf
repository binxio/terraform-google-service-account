locals {
  project     = var.project
  environment = "all"

  service_accounts = {
    "trigger-assertions for sa 'cause this name is too long and has invalid chars" = {}
  }
}

module "sa" {
  source = "../../"

  project     = local.project
  environment = local.environment

  service_accounts = local.service_accounts
}
