resource "aws_db_instance" "main" {
#   identifier              = "${local.prefix}-db"
#   db_name                 = "weshare"
  name = "weshare2"
  identifier = "weshare2"
  allocated_storage       = 10
  storage_type            = "gp2"
  engine                  = "postgres"
  engine_version          = "14.6"
  instance_class          = "db.t3.micro"
#   db_subnet_group_name    = aws_db_subnet_group.main.name
  password                = var.db_password
  username                = var.db_username
  backup_retention_period = 0
  multi_az                = false
  skip_final_snapshot      = true
  publicly_accessible = true
#   vpc_security_group_ids  = [aws_security_group.rds.id]
  port = 5432
#   tags = merge(
#     local.common_tags,
#     map("Name", "${local.prefix}-main")
#   )
}
# locals {
#   prefix = "${var.prefix}-${terraform.workspace}"
#   common_tags = {
#     Environment = terraform.workspace
#     Project     = var.project
#     Owner       = var.contact
#     ManagedBy   = "Terraform"
#   }
# }

