terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.54.1"
    }
  }
}


resource "aws_security_group" "rds_sg" {
  name        = "rds-mysql-sg"
  description = "Allow MySQL inbound traffic"

  ingress {
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}





resource "aws_db_instance" "default"{
  vpc_security_group_ids = [aws_security_group.rds_sg.id]
  engine = "mysql"
  instance_class = "db.t4g.micro"
  db_name = "url_mappings"
  allocated_storage = "5"
  engine_version = "8.0.39"
  skip_final_snapshot = "true"
  username = "url_shortener"
  password = "fortheapp"
  publicly_accessible = true
}