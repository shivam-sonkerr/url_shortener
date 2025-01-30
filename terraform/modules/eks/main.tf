# variable "security_group_ids" {
#   description = "Security group IDs to attach to EKS worker nodes"
#   type        = list(string)
# }



resource "aws_eks_cluster" "main" {
  name     = "url-shortener"
  role_arn = var.eks_role_arn

  vpc_config {
    subnet_ids = var.subnet_ids

  }
}

resource "aws_eks_node_group" "workers" {
  cluster_name    = aws_eks_cluster.main.name
  node_group_name = "worker-nodes"
  disk_size = 20
  node_role_arn   = "arn:aws:iam::339713031726:role/worker_nodes_role"  # Replace with an IAM role ARN
  subnet_ids = var.private_subnet_ids


  scaling_config {
    desired_size = 1
    min_size     = 1
    max_size     = 2
  }

  instance_types = ["t3.medium"]  # Use a basic instance type
}
