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


resource "aws_security_group" "eks_worker_sg" {
  name        = "eks-worker-nodes-sg"
  description = "Allow inbound traffic to EKS worker nodes"
  vpc_id      = var.vpc_id

  ingress {
    from_port   = 1025
    to_port     = 65535
    protocol    = "tcp"
    cidr_blocks = [var.vpc_cidr_block]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}



output "worker_sg_id" {
  description = "Security Group ID for EKS worker nodes"
  value       = aws_security_group.eks_worker_sg.id
}


