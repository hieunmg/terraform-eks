# module "ec2" {
#   source = "./modules/network/ec2"
#   subnet_id = module.subnet.public_subnet_id
#   instance_type = var.instance_type
#   instance_name = var.instance_name
# }

module "vpc" {
  source = "./modules/network/vpc"
  my_az_01 = var.my_az_01
  my_az_02 = var.my_az_02
  igw_name = var.igw_name
  rtb_name = var.rtb_name
  public_subnet_name_01 = var.public_subnet_name_01
  public_subnet_name_02 = var.public_subnet_name_02
  public_subnet_01_cidr = var.public_subnet_01_cidr
  public_subnet_02_cidr = var.public_subnet_02_cidr
}

# module "subnet" {
#   source = "./modules/network/subnet"
#   vpc_id = module.vpc.vpc_id
# }

module "eks" {
  source = "./modules/eks"
  eks_role_arn = module.iam.eks_role_arn
  eks_nodegroup_arn = module.iam.eks_nodegroup_arn
  eks_cluster_name = var.eks_cluster_name
  my_public_subnet_01_id = module.vpc.my_public_subnet_01_id
  my_public_subnet_02_id = module.vpc.my_public_subnet_02_id
}

module "iam" {
  source = "./modules/iam"
}
