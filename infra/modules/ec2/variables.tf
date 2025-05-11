variable "vpc_id" {}
variable "subnet_id" {}
variable "ami_id" {
  default = "ami-04e7764922e1e3a57"
}
variable "instance_type" {
  default = "t2.micro"
}
variable "key_name" {}
variable "my_ip" {}
variable "environment" {}
variable "repo_url" {}
variable "server_port" {
  default = 8080
}
