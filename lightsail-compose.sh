#!/bin/bash

# Install the latest version of Docker the lazy way
curl -sSL https://get.docker.com | sh

# Make it so you don't need to sudo to run Docker commands
usermod -aG docker ubuntu
# Install Docker Compose
curl -L https://github.com/docker/compose/releases/download/1.21.2/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

# Install Git
apt-get update
apt-get install -y git

# Configure Git
git config --global core.autocrlf false

# Copy the Dockerfile into /srv/docker 
# If you change this, change the systemd service file to match
# WorkingDirectory=[whatever you have below]
mkdir /srv/docker
curl -o /srv/docker/docker-compose.yml https://raw.githubusercontent.com/aidostt/diploma-backend-deploy/main/docker-compose.yml

# Copy in systemd unit file and register it so our compose file runs 
# on system restart
curl -o /etc/systemd/system/docker-compose-app.service https://raw.githubusercontent.com/aidostt/diploma-backend-deploy/main/docker-compose-app.service
systemctl enable docker-compose-app

# Start up the application via Docker Compose
docker-compose -f /srv/docker/docker-compose.yml up -d

