variable "bucket_name" {
  description = "The name of the S3 bucket."
  type        = string
}

variable "environment" {
  description = "Deployment environment"
  type        = string
}

variable "index_file_path" {
  description = "Local path to index.html for static site."
  type        = string
}
variable "api_public_dns" {
  description = "Public DNS of the EC2 instance for injecting into client"
  type        = string
}
