provider "aws" {
  region = var.aws_region
}

resource "aws_security_group" "rds_sg" {
  vpc_id = var.vpc_id
  name        = "rds-mysql-sg"
  description = "Allow MySQL inbound traffic from EKS nodes"

  ingress {
    from_port       = 3306
    to_port         = 3306
    protocol        = "tcp"
    security_groups = [var.eks_worker_sg_id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "rds-security-group"
  }
}


resource "aws_db_subnet_group" "default" {
  name       = "private-db-subnet-group-${random_id.suffix.hex}"
  subnet_ids = var.subnet_ids

  tags = {
    Name = "private-db-subnet-group"
  }
}

resource "random_id" "suffix" {
  byte_length = 8
}


resource "aws_db_instance" "default" {
  vpc_security_group_ids = [aws_security_group.rds_sg.id]
  engine                 = "mysql"
  instance_class         = "db.t4g.micro"
  db_name                = var.db_name
  allocated_storage      = var.allocated_storage
  engine_version = "8.0.39"
  skip_final_snapshot    = true
  username               = var.db_username
  password               = var.db_password
  publicly_accessible    = false
  db_subnet_group_name  = aws_db_subnet_group.default.name

  tags = {
    Name = "private-rds-instance"
  }
}



