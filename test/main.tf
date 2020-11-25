locals {
  project     = var.project
  environment = "all"
  test_sas = {
    "user" = format("sa-user-%s", replace(local.environment, " ", "-"))
  }
}

resource "google_service_account" "map" {
  for_each = local.test_sas

  account_id   = each.value
  display_name = format("%s Terraform Test", each.value)
  description  = "Service Account to test assignment of service account roles"
}

