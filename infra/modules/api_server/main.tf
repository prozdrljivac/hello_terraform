resource "aws_security_group" "hello_terraform_server_sg" {
  name        = "${var.environment}-hello-terraform-server-sg"
  description = "Allow HTTP and SSH"
  vpc_id      = var.vpc_id

  ingress {
    description = "Allow HTTP"
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "Allow SSH"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = [var.my_ip]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "hello_terraform_server" {
  ami                    = var.ami_id
  instance_type          = var.instance_type
  key_name               = var.key_name
  vpc_security_group_ids = [aws_security_group.hello_terraform_server_sg.id]
  subnet_id              = var.subnet_id

  user_data = templatefile("${path.module}/startup.sh.tpl", {
    repo_url    = var.repo_url
    server_port = var.server_port
  })

  tags = {
    Name = "hello-terraform-server-${var.environment}"
  }
}
