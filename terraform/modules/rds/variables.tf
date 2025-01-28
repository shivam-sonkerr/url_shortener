variable "cidr_block" {
  description = "The CIDR block for the VPC"
  type        = string
  default     = "10.0.0.0/16"
}

variable "private_subnet_count" {
  description = "Number of private subnets"
  type        = number
  default     = 2
}

variable "private_subnet_cidrs" {
  description = "CIDR blocks for private subnets"
  type        = list(string)
  default     = ["10.0.1.0/24", "10.0.2.0/24"]
}

variable "public_subnet_count" {
  description = "Number of public subnets"
  type        = number
  default     = 2
}

variable "public_subnet_cidrs" {
  description = "CIDR blocks for public subnets"
  type        = list(string)
  default     = ["10.0.3.0/24", "10.0.4.0/24"]
}



variable "db_name" {
  description = "The name of the database"
  type        = string
  default     = "url_mappings"
}

variable "db_username" {
  description = "The database username"
  type        = string
  default     = "url_shortener"
}

variable "db_password" {
  description = "The database password"
  type        = string
  default     = "fortheapp"
}

variable "allocated_storage" {
  description = "The allocated storage for the RDS instance"
  type        = number
  default     = 5
}
