variable "bucket_name" {
  description = "Name of the S3 bucket that stores static-site assets."
  type        = string
}

variable "environment" {
  description = "Deployment environment (e.g. dev, staging, prod)."
  type        = string
}

variable "api_public_dns" {
  description = "Public DNS hostname of the API EC2 instance, injected into the client."
  type        = string
}
