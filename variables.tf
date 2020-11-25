#------------------------------------------------------------------------------------------------------------------------
# 
# Generic variables
#
#------------------------------------------------------------------------------------------------------------------------
variable "project" {
  description = "Company project name."
  type        = string
}

variable "environment" {
  description = "Company environment for which the resources are created (e.g. dev, tst, acc, prd, all)."
  type        = string
}

variable "gcp_project" {
  description = "GCP Project ID override - this is normally not needed and should only be used in tf-projects."
  type        = string
  default     = null
}

#------------------------------------------------------------------------------------------------------------------------
#
# Service account variables
#
#------------------------------------------------------------------------------------------------------------------------

variable "service_accounts" {
  description = "Map of accounts to be created. The key will be used for the SA name so it should describe the SA purpose. A list of roles can be provided to set an authoritive iam policy for the service account."

  # Don't define type as it will require roles to be provided
  #type = object({
  #  roles = map(map(string))
  #})
}
