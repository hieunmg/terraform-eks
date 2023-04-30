
## Terraform block
terraform {
  required_version = "~> 1.3.4"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.48.0"
    }
  }

  # backend "s3" {
  #   bucket         = "weshare-terraform-s3-backend"
  #   key            = "weshare-app.tfstate"
  #   region         = "ap-southeast-1"
  #   encrypt        = true
  #   dynamodb_table = "weshare-terraform-tf-state-lock"
  # }
}

## Provider block
provider "aws" {
  region  = "ap-southeast-1"
  profile = "default"
}
