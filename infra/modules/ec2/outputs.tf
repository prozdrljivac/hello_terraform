output "instance_public_ip" {
  value = aws_instance.hello_terraform_server.public_ip
}
