# Secrets
variable "access_key" {
  description = "AWS access key for main account"
  type        = string
  sensitive   = true
}

variable "secret_key" {
  description = "AWS secret key key for main account"
  type        = string
  sensitive   = true
}

# Non sensible variables
variable "region" {
  description = "AWS region for the deployment of the resources"
  type        = string
  sensitive   = false
}

variable "availability_zone" {
  description = "AWS availability zone for the deployment of the subnet and instance"
  type        = string
  sensitive   = false
}

