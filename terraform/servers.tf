data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

resource "aws_key_pair" "dev" {
  key_name   = "dev-key"
  public_key = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIMYxDwgna/0mFFaQCfTPXyfiItQWANftAu+jUPbw1s1A"

  tags = {
    App = "nginx-random-redirect"
  }
}

resource "aws_instance" "web" {
  ami             = data.aws_ami.ubuntu.id
  instance_type   = "t2.nano"
  security_groups = ["${aws_security_group.ingress-ssh.name}"]

  key_name = aws_key_pair.dev.key_name

  tags = {
    App = "nginx-random-redirect"
  }
}

output "public_dns" {
  value = aws_instance.web.public_dns
}
