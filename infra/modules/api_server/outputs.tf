output "instance_public_ip" {
  description = "Public IPv4 address of the API EC2 instance."
  value       = aws_instance.hello_terraform_server.public_ip
}

output "instance_public_dns" {
  description = "Public DNS hostname of the API EC2 instance."
  value       = aws_instance.hello_terraform_server.public_dns
}
