output "bucket_name" {
  value = aws_s3_bucket.static_site_bucket.id
}

output "website_url" {
  value = aws_s3_bucket_website_configuration.static_site_configuration.website_endpoint
}
