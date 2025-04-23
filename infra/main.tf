resource "aws_instance" "api" {
  ami           = var.api_ami_id
  instance_type = var.api_instance_type
}
