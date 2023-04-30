resource "aws_vpc" "myvpc" {
  cidr_block = var.vpc_cidr_block
  tags = {
    Name = var.vpc_name
  }
}

resource "aws_subnet" "my_public_subnet_01" {
    vpc_id = aws_vpc.myvpc.id
    availability_zone = var.my_az_01
    cidr_block = var.public_subnet_02_cidr
    map_public_ip_on_launch = true
    tags = {
      Name = var.public_subnet_name_01
    }
}

resource "aws_subnet" "my_public_subnet_02" {
    availability_zone = var.my_az_02
    vpc_id = aws_vpc.myvpc.id
    cidr_block = var.public_subnet_01_cidr
    map_public_ip_on_launch = true
    tags = {
      Name = var.public_subnet_name_02
    }
}


resource "aws_security_group" "mysg" {
    name = var.sg_name
    vpc_id = aws_vpc.myvpc.id

    tags = {
      Name = var.sg_name
    }

    dynamic "ingress" {
        iterator = port
        for_each = var.ingress
        content {
          from_port = port.value
          to_port = port.value
          protocol = "TCP"
          cidr_blocks = [ "0.0.0.0/0" ]
        }
    }
    egress {
          from_port = 0
          to_port = 0
          protocol = "-1"
          cidr_blocks = [ "0.0.0.0/0" ]
    }    
}


resource "aws_route_table" "myroutetable" {
    vpc_id = aws_vpc.myvpc.id
    route {
        cidr_block = "0.0.0.0/0"
        gateway_id = aws_internet_gateway.myigw.id
    }
    
    tags = {
      Name = var.rtb_name
    }
}

resource "aws_internet_gateway" "myigw" {
    vpc_id = aws_vpc.myvpc.id
    tags = {
      "Name" = var.igw_name
    }
  
}
resource "aws_route_table_association" "publicsubnet_asso1" {
    subnet_id = aws_subnet.my_public_subnet_01.id
    route_table_id = aws_route_table.myroutetable.id
}
resource "aws_route_table_association" "publicsubnet_asso2" {
    subnet_id = aws_subnet.my_public_subnet_02.id
    route_table_id = aws_route_table.myroutetable.id
}

