module "vpc" {
  source = "../vpc"  # Adjust the path according to your folder structure
  cidr_block = var.cidr_block
  private_subnet_count = var.private_subnet_count
  private_subnet_cidrs = var.private_subnet_cidrs
  public_subnet_count = var.public_subnet_count
  public_subnet_cidrs = var.public_subnet_cidrs
}


resource "aws_eks_cluster" "main" {
  name     = "url-shortener"
  role_arn = var.eks_role_arn

  vpc_config {
    subnet_ids = module.vpc.private_subnet_ids
    security_group_ids = []
  }
}

resource "aws_eks_node_group" "workers" {
  cluster_name    = aws_eks_cluster.main.name
  node_group_name = "worker-nodes"
  disk_size = 5
  node_role_arn   = "arn:aws:iam::339713031726:role/worker_nodes_role"  # Replace with an IAM role ARN
  subnet_ids      = module.vpc.private_subnet_ids

  scaling_config {
    desired_size = 1
    min_size     = 1
    max_size     = 2
  }

  instance_types = ["t2.micro"]  # Use a basic instance type
}
