variable "project" {
  description = "Project name passed as var so we can randomize the resource id"
  type        = string
}
variable "sa_user_email" {
  description = "The User SA used for testing service account role assignment"
  type        = string
  default     = ""
}
