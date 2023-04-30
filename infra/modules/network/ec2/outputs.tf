
output "ec2" {
  value = aws_instance.ansible_server.public_ip
}