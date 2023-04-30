variable "instance_type" {
  description = "Instance type"
  type        = string
  # default     = "t2.micro"
}

variable "instance_name" {
  description = "Instance name"
  type = string
  default = "Ansible Server"
}

variable "subnet_id" {
  description = "Subnet id"
  type = string
}