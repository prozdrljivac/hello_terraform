output "bucket_name" {
  description = "Name of the S3 bucket hosting the static site."
  value       = aws_s3_bucket.static_site_bucket.id
}

output "website_url" {
  description = "Website endpoint URL for the static site."
  value       = aws_s3_bucket_website_configuration.static_site_configuration.website_endpoint
}
