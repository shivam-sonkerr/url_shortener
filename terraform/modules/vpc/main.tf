# SHARED VPC BETWEEN RDS AND EKS

resource "aws_vpc" "main" {
  cidr_block = var.cidr_block
  enable_dns_support = true
  enable_dns_hostnames = true

  tags = {
    Name = "shared-vpc"
  }
}



resource "aws_subnet" "private" {
  count = var.private_subnet_count
  vpc_id = aws_vpc.main.id
  cidr_block = element(var.private_subnet_cidrs,count.index)
  availability_zone = element(["us-west-2a", "us-west-2b"], count.index)

  tags = {
    Name = "private-subnet-${count.index}"
  }
}


resource "aws_subnet" "public" {
  count = var.public_subnet_count
  vpc_id = aws_vpc.main.id
  cidr_block = element(var.public_subnet_cidrs,count.index )
  availability_zone = element(["us-west-2a", "us-west-2b"], count.index)

  tags = {
    Name = "public-subnet-${count.index}"
  }
}


# Create a NAT Gateway in one of the public subnets
resource "aws_eip" "nat" {
  count = var.private_subnet_count
  vpc   = true
}

resource "aws_nat_gateway" "nat" {
  count         = var.private_subnet_count
  allocation_id = element(aws_eip.nat[*].id, count.index)
  subnet_id     = element(aws_subnet.public[*].id, count.index) # One NAT per AZ

  tags = {
    Name = "nat-gateway-${count.index}"
  }
}

resource "aws_route_table" "private" {
  count = var.private_subnet_count
  vpc_id = aws_vpc.main.id

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = element(aws_nat_gateway.nat[*].id, count.index)
  }

  tags = {
    Name = "private-route-table-${count.index}"
  }
}

resource "aws_route_table_association" "private" {
  count          = var.private_subnet_count
  subnet_id      = element(aws_subnet.private[*].id, count.index)
  route_table_id = element(aws_route_table.private[*].id, count.index)
}



# Internet Gateway
resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "internet-gateway"
  }
}

# Public Route Table
resource "aws_route_table" "public" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.main.id
  }

  tags = {
    Name = "public-route-table"
  }
}

# Associate Public Subnets with the Public Route Table
resource "aws_route_table_association" "public" {
  count          = length(aws_subnet.public)
  subnet_id      = element(aws_subnet.public[*].id, count.index)
  route_table_id = aws_route_table.public.id
}


