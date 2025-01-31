provider "aws" {
  region = "us-west-2"  # Adjust the region as needed
}

resource "aws_security_group" "eks_worker_sg" {
  vpc_id = module.vpc.vpc_id  # Reference to the VPC ID

  egress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
  }

  ingress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
  }

  tags = {
    Name = "eks-worker-sg"
  }
}




# VPC Module
module "vpc" {
  source               = "./modules/vpc"
  cidr_block           = "10.0.0.0/16"  # Adjust the CIDR as needed
  private_subnet_count = 2  # 2 private subnets across different AZs
  public_subnet_count  = 2  # 2 public subnets across different AZs
  private_subnet_cidrs = ["10.0.1.0/24", "10.0.2.0/24"]  # 2 subnets in different AZs
  public_subnet_cidrs  = ["10.0.3.0/24", "10.0.4.0/24"]  # 2 subnets in different AZs
  private_subnet_ids = module.vpc.private_subnet_ids
}

# EKS Module
module "eks" {
  source             = "./modules/eks"
  eks_role_arn       = "arn:aws:iam::339713031726:role/eks-cluster-role"
  subnet_ids = module.vpc.private_subnet_ids
  private_subnet_ids = module.vpc.private_subnet_ids
}


# RDS Module
module "rds" {
  source            = "./modules/rds"
  db_name           = "url_mappings"
  db_username       = "url_shortener"
  db_password       = "fortheapp"
  allocated_storage = 20
  subnet_ids        = module.vpc.private_subnet_ids  # Ensure this line references the output
}



