
variable "vpc_name" {
  description = "VPC Name"
  type        = string
  default     = "myvpc"
}

variable "vpc_cidr_block" {
  description = "VPC CIDR Block"
  type        = string
  default     = "10.0.0.0/16"
}

variable "vpc_availability_zones" {
  description = "VPC Availability Zones"
  type        = list(string)
  default     = ["ap-southeast-1a"]
}

variable "instance_type" {
  description = "Instance type"
  type        = string
  default     = "t2.micro"

  validation {
    condition = contains(["t2.micro", "t3.small", "t3.micro"], var.instance_type)
    error_message = "Value not allow."
  }
}

variable "instance_name" {
  description = "Instance name"
  type        = string
  default = "myec2"
}

variable "eks_version" {
  default = "1.25"
}

variable "project" {
  description = "The project name to use for unique resource naming"
  default     = "weshare-terraform"
  type        = string
}
variable "principal_arns" {
  description = "A list of principal arns allowed to assume the IAM role"
  default     = null
  type        = list(string)
}

variable "eks_cluster_name" {
  type = string
  default = "weshare"
}

variable "sg_name" {
  type = string
  default = "mysg"
}

variable "igw_name" {
  type = string
  default = "myigw"
}

variable "ingress" {
    type = list(number)
    default = [ 22, 80, 443 ]
}

variable "public_subnet_name_01" {
  type        = string
  default = "my_public_subnet_01"
}
variable "public_subnet_name_02" {
  type        = string
  default = "my_public_subnet_02"
}


variable "rtb_name" {
  type = string
  default = "myrtb"
}


variable "my_az_01" {
  description = "Availability Zones"
  type = string
}

variable "my_az_02" {
  description = "Availability Zones"
  type = string
}

variable "public_subnet_01_cidr" {
  type = string
}

variable "public_subnet_02_cidr" {
  type = string
}