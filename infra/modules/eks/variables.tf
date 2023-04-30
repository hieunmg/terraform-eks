variable "eks_cluster_name" {
  type = string
  default = "weshare"
}

variable "instance_name" {
  default = "myec2"
}

variable "eks_version" {
  type = string
  default = "1.25"
}

variable "eks_node_group_instance_types" {
  type = list(string)
  default = ["t3.small"]
}

variable "endpoint_public_access" {
  type = bool
  default = true
}

variable "endpoint_private_access" {
  type = bool
  default = true
}

variable "node_group_name" {
  type = string
  default = "weshare-nodegroup"
}

variable "node_role_arn" {
  type = string
  default = "arn:aws:iam::023221543387:role/weshare-node"
}

variable "eks_role_arn" {
  description = "Cluster role resource name"
  type = string
}

variable "eks_nodegroup_arn" {
  description = "Node group role resource name"
  type = string
}

variable "my_public_subnet_01_id" {
  type = string
}
variable "my_public_subnet_02_id" {
  type = string
}

