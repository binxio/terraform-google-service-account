locals {
  project     = var.project
  environment = var.environment

  # Check if we want to deploy to a specific gcp_project or not.
  gcp_project = (var.gcp_project == null ? null : (length(var.gcp_project) > 0 ? var.gcp_project : null))

  service_accounts = { for sa, settings in var.service_accounts : sa =>
    {
      account_id   = replace(sa, "/[^\\p{Ll}\\p{Lo}\\p{N}-]+/", "-")
      display_name = title(format("Terraform Managed SA: %s %s %s", local.project, local.environment, sa))
      roles        = { for role, members in try(settings.roles, {}) : role => [for member, type in members : format("%s:%s", type, member)] }
    }
  }
}

resource "google_service_account" "map" {
  project  = local.gcp_project
  for_each = local.service_accounts

  account_id   = each.value.account_id
  display_name = each.value.display_name
}

data "google_iam_policy" "map" {
  for_each = { for service_account, settings in local.service_accounts : service_account => settings if settings.roles != null }

  dynamic "binding" {
    for_each = each.value.roles

    content {
      role    = binding.key
      members = binding.value
    }
  }
}

resource "google_service_account_iam_policy" "map" {
  for_each = data.google_iam_policy.map

  service_account_id = google_service_account.map[each.key].name
  policy_data        = each.value.policy_data
}
