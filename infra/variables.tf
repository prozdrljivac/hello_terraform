variable "aws_region" {
  description = "AWS region to deploy to"
  type        = string
}

variable "aws_access_key" {
  description = "AWS access key"
  type        = string
  sensitive   = true
}

variable "aws_secret_key" {
  description = "AWS secret key"
  type        = string
  sensitive   = true
}

variable "api_ami_id" {
  description = "AMI ID for the Go API EC2 instance"
  type        = string
}

variable "api_instance_type" {
  description = "EC2 instance type for API"
  type        = string
}
