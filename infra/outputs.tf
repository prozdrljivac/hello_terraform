output "api_public_ip" {
  description = "Public IPv4 address of the API EC2 instance exposed by the api_server module."
  value       = module.api_server.instance_public_ip
}

output "api_public_dns" {
  description = "Public DNS hostname of the API EC2 instance."
  value       = module.api_server.instance_public_dns
}

output "client_url" {
  description = "Base URL of the client-facing static website served by the static_website module."
  value       = module.static_website.website_url
}

output "static_site_bucket_name" {
  description = "Name of the S3 bucket that stores the static website files."
  value       = module.static_website.bucket_name
}
