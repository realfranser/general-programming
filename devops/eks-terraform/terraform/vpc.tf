resource "aws_vpc" "main_vpc" {
  cidr_block = "192.168.0.0/16"

  # Makes your instances shared on the host
  instance_tenancy = "default"

  # Required by EKS. Enable/disable DNS support in the VPC
  enable_dns_support = true
  # Required by EKS. Enable/disable DNS hostnames in the VPC
  enable_dns_hostnames = true

  # Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length
  assign_generated_ipv6_cidr_block = false

  tags = {
    Name    = "main vpc"
    Project = "eks-terraform"
    Repo    = "https://github.com/realfranser/general-programming/tree/main/devops/eks-terraform"
  }
}

output "vpc_id" {
  value       = aws_vpc.main_vpc.id
  description = "Main VPC id"
}
