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
  # Amazon Linux 2
  default = "ami-04e7764922e1e3a57"
}

variable "api_instance_type" {
  description = "EC2 instance type for API"
  type        = string
  default     = "t2.micro"
}

variable "bucket_name" {
  description = "The name for the S3 bucket."
  type        = string
}

variable "environment" {
  description = "Deployment environment"
  type        = string
  default     = "dev"
}

variable "index_file_path" {
  description = "Path to the index.html file to upload."
  type        = string
}

variable "vpc_id" {}
variable "subnet_id" {}
variable "key_name" {
  description = "Name of the existing EC2 key pair for SSH access"
  type        = string
}

variable "my_ip" {
  description = "Your IP to allow SSH access to EC2"
  type        = string
}

variable "repo_url" {}
variable "server_port" {
  description = "TODO"
  type        = string
  default     = 8080
}
