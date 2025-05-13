#!/bin/bash

# Update and install dependencies
yum update -y
yum install -y git wget

# Install Go
cd /usr/local
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz

# Set PATH for ec2-user permanently
echo 'export PATH=$PATH:/usr/local/go/bin' >> /home/ec2-user/.bashrc

# Clone repo
cd /home/ec2-user
git clone ${repo_url} app
chown -R ec2-user:ec2-user ./app

# Create .env file
cat > /home/ec2-user/app/api/.env <<EOF
SERVER_PORT=${server_port}
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=hello_terraform
DB_PORT=5432
DB_HOST=localhost
EOF

chown ec2-user:ec2-user /home/ec2-user/app/api/.env

# Create systemd service
cat > /etc/systemd/system/go-api.service <<EOF
[Unit]
Description=Go API Server
After=network.target

[Service]
User=ec2-user
WorkingDirectory=/home/ec2-user/app/api
EnvironmentFile=/home/ec2-user/app/api/.env
ExecStart=/usr/local/go/bin/go run ./cmd/api-server/app.go
Restart=on-failure
RestartSec=5
StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target
EOF

# Enable and start the service
systemctl daemon-reexec
systemctl daemon-reload
systemctl enable go-api
systemctl start go-api
