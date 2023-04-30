
output "vpc_id" {
    value = aws_vpc.myvpc.id
}

output "my_public_subnet_01_id" {
    value = aws_subnet.my_public_subnet_01.id
}
output "my_public_subnet_02_id" {
    value = aws_subnet.my_public_subnet_02.id 
}
output "my_sg_id" {
    value = aws_security_group.mysg.id
}
