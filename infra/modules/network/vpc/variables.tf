variable "vpc_name" {
  description = "VPC Name"
  type        = string
  default     = "myvpc"
}

variable "vpc_cidr_block" {
  description = "VPC CIDR Block"
  type        = string
  default     = "10.10.0.0/16"
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