variable "vpc_id" {
  description = "ID of the VPC where the API server will be deployed."
  type        = string
}

variable "subnet_id" {
  description = "ID of the public subnet for the API server."
  type        = string
}

variable "ami_id" {
  description = "AMI ID for the API EC2 instance."
  type        = string
  # Amazon Linux 2023
  default = "ami-04e7764922e1e3a57"
}

variable "instance_type" {
  description = "EC2 instance size for the API server."
  type        = string
  default     = "t2.micro"
}

variable "key_name" {
  description = "Name of an existing EC2 key pair for SSH access."
  type        = string
}

variable "my_ip" {
  description = "Your public IPv4 address to allow SSH access from."
  type        = string
}

variable "environment" {
  description = "Deployment environment (e.g. dev, staging, prod)."
  type        = string
}

variable "repo_url" {
  description = "Git repository URL containing the API source code."
  type        = string
}

variable "server_port" {
  description = "Port on which the API listens inside the EC2 instance."
  type        = number
  default     = 8080
}
