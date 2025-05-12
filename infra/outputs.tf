output "api_public_ip" {
  value = module.ec2_backend.instance_public_ip
}
output "api_public_dns" {
  value = module.ec2_backend.instance_public_dns
}

output "client_url" {
  value = module.s3_static_site.website_url
}

output "static_site_bucket_name" {
  value = module.s3_static_site.bucket_name
}
