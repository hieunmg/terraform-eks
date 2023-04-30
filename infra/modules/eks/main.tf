
resource "aws_eks_cluster" "cluster" {
    name = var.eks_cluster_name
    vpc_config {
        subnet_ids = [var.my_public_subnet_01_id, var.my_public_subnet_02_id]
        endpoint_public_access = var.endpoint_public_access
        endpoint_private_access = var.endpoint_private_access
        # vpc_id = aws_default_vpc.default.id
    }

    role_arn = var.eks_role_arn
  #   depends_on = [
  #   aws_iam_role_policy_attachment.example-AmazonEKSClusterPolicy,
  #   aws_iam_role_policy_attachment.example-AmazonEKSVPCResourceController,
  # ]
}
resource "aws_eks_node_group" "eks_nodegroup" {
  cluster_name = aws_eks_cluster.cluster.name
  version = aws_eks_cluster.cluster.version
  node_group_name = var.node_group_name
  node_role_arn   = var.eks_nodegroup_arn
  subnet_ids      = [var.my_public_subnet_01_id, var.my_public_subnet_02_id]
  instance_types = var.eks_node_group_instance_types
  scaling_config {
    desired_size = 1
    max_size     = 2
    min_size     = 1
  }
  update_config {
    max_unavailable = 1
  }
  depends_on = [aws_eks_cluster.cluster]
}