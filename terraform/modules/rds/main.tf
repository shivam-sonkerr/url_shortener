provider "aws" {
  region = "us-west-2"  # Adjust the region as needed
}





resource "aws_security_group" "rds_sg" {
  name        = "rds-mysql-sg"
  description = "Allow MySQL inbound traffic from EKS nodes"

  ingress {
    from_port       = 3306
    to_port         = 3306
    protocol        = "tcp"
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
  name       = "private-db-subnet-group"
  subnet_ids = var.subnet_ids

  tags = {
    Name = "private-db-subnet-group"
  }
}

resource "aws_db_instance" "default" {
  vpc_security_group_ids = [aws_security_group.rds_sg.id]
  engine                 = "mysql"
  instance_class         = "db.t3.medium"
  db_name                = var.db_name
  allocated_storage      = var.allocated_storage
  engine_version         = "8.0.39"
  skip_final_snapshot    = true
  username               = var.db_username
  password               = var.db_password
  publicly_accessible    = false

  tags = {
    Name = "private-rds-instance"
  }
}
