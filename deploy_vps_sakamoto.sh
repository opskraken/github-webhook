#!/bin/bash
set -e

REPO_DIR=/app/vps-sakamoto

# Clone or update repo
if [ ! -d "$REPO_DIR/.git" ]; then
    git clone git@github.com:opskraken/vps-sakamoto.git $REPO_DIR
else
    cd $REPO_DIR
    git reset --hard
    git pull origin main
fi

# Rebuild and redeploy Docker containers
docker-compose -f $REPO_DIR/docker-compose.yaml down
docker-compose -f $REPO_DIR/docker-compose.yaml up -d --build
