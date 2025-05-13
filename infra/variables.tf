variable "aws_region" {
  description = "AWS region where resources will be deployed."
  type        = string
}

variable "aws_access_key" {
  description = "Access-key ID for the dedicated Terraform user."
  type        = string
  sensitive   = true
}

variable "aws_secret_key" {
  description = "Secret access key for the dedicated Terraform user."
  type        = string
  sensitive   = true
}

variable "api_ami_id" {
  description = "AMI ID for the EC2 instance that runs the API."
  type        = string
  # Amazon Linux 2023
  default = "ami-04e7764922e1e3a57"
}

variable "api_instance_type" {
  description = "EC2 instance type for the API server."
  type        = string
  default     = "t2.micro"
}

variable "bucket_name" {
  description = "Name of the S3 bucket that stores static assets."
  type        = string
}

variable "environment" {
  description = "Deployment environment (e.g. dev, staging, prod)."
  type        = string
  default     = "dev"
}

variable "vpc_id" {
  description = "ID of the VPC where resources will be created."
  type        = string
}

variable "subnet_id" {
  description = "ID of the public subnet for the EC2 instance."
  type        = string
}

variable "key_name" {
  description = "Name of an existing EC2 key pair for SSH access."
  type        = string
}

variable "my_ip" {
  description = "Your public IPv4 address to allow SSH access from."
  type        = string
}

variable "repo_url" {
  description = "Git repository URL containing the application source."
  type        = string
}

variable "server_port" {
  description = "Port on which the application listens inside the EC2 instance."
  type        = number
  default     = 8080
}
