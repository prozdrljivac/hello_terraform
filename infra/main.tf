# Todo - Add Terraform lock on AWS version

# Discover default VPC
data "aws_vpc" "default" {
  default = true
}

# Use first public subnet in the default VPC
data "aws_subnets" "default" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.default.id]
  }
}
data "aws_subnet" "default" {
  id = tolist(data.aws_subnets.default.ids)[0]
}
module "s3_static_site" {
  source          = "./modules/s3"
  bucket_name     = var.bucket_name
  environment     = var.environment
  index_file_path = "" # not used anymore
  api_public_dns  = module.ec2_backend.instance_public_dns
}

module "ec2_backend" {
  source      = "./modules/ec2"
  environment = var.environment
  vpc_id      = data.aws_vpc.default.id
  subnet_id   = data.aws_subnet.default.id
  key_name    = var.key_name
  my_ip       = var.my_ip
  repo_url    = "https://github.com/prozdrljivac/hello_terraform.git"
  server_port = 8080
}
