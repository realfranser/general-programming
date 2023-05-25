resource "aws_vpc" "main_vpc" {
  cidr_block = "192.168.0.0/16"

  # Makes your instances shared on the host
  instance_tenancy = "default"

  # Required by EKS. Enable/disable DNS support in the VPC
  enable_dns_support = true
  # Required by EKS. Enable/disable DNS hostnames in the VPC
  enable_dns_hostnames = true

  # Enable/disable ClassicLink for the VPC
  enable_classiclink = false
  # Enable/disable ClassicLink DNS support for the VPC
  enable_classiclink_dns_support = false

  # Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length
  assign_generated_ipv6_cidr_block = false

  tags = {
    Name = "main vpc"
  }
}

output "vpc_id" {
  value       = aws_vpc.main_vpc.id
  description = "Main VPC id"
}
