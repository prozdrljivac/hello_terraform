terraform {
  required_version = "~> 1.11.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.96.0"
    }
  }
}

data "aws_vpc" "default" {
  default = true
}

data "aws_subnets" "default" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.default.id]
  }
}
data "aws_subnet" "default" {
  id = tolist(data.aws_subnets.default.ids)[0]
}
module "static_website" {
  source         = "./modules/static_website"
  bucket_name    = var.bucket_name
  environment    = var.environment
  api_public_dns = module.api_server.instance_public_dns
}

module "api_server" {
  source      = "./modules/api_server"
  environment = var.environment
  vpc_id      = data.aws_vpc.default.id
  subnet_id   = data.aws_subnet.default.id
  key_name    = var.key_name
  my_ip       = var.my_ip
  repo_url    = "https://github.com/prozdrljivac/hello_terraform.git"
  server_port = 8080
}
