#!/bin/bash
set -e

echo "====================================================================="
echo " ComplyAge Dockerized Deployment Script "
echo "====================================================================="

REPO_DIR="/complyage/complyage.com"
ARIA_DIR="$REPO_DIR/aria"
COMPOSE_FILE="$REPO_DIR/prod.docker-compose.yml"
NGINX_CONF_SRC="$REPO_DIR/.nginx/compose-proxy.conf"
NGINX_CONF_DST="/etc/nginx/conf.d/compose-proxy.conf"

#------------------------------------------------------------------------||
#-|| [1/8] Stop running containers
#------------------------------------------------------------------------||

echo "[1/8] Stopping any existing Docker containers..."
docker compose -f "$COMPOSE_FILE" down || true

#------------------------------------------------------------------------||
#-|| [2/8] Update or clone Aria Framework
#------------------------------------------------------------------------||

echo "[2/8] Updating Aria Framework..."
if [ -d "$ARIA_DIR/.git" ]; then
    echo "Aria already exists — pulling latest changes..."
    cd "$ARIA_DIR" && git reset --hard && git pull origin main
else
    echo "Aria not found — cloning fresh..."
    git clone https://github.com/ralphferrara/aria.git "$ARIA_DIR"
fi

#------------------------------------------------------------------------||
#-|| [3/8] Clean and rebuild Docker images
#------------------------------------------------------------------------||

echo "[3/8] Rebuilding all containers..."
cd "$REPO_DIR"
docker compose -f "$COMPOSE_FILE" build --no-cache

#------------------------------------------------------------------------||
#-|| [4/8] Start Docker stack
#------------------------------------------------------------------------||

echo "[4/8] Starting Docker stack..."
docker compose -f "$COMPOSE_FILE" up -d

#------------------------------------------------------------------------||
#-|| [5/8] Apply updated Nginx configuration
#------------------------------------------------------------------------||

echo "[5/8] Updating Nginx reverse proxy configuration..."
cp -f "$NGINX_CONF_SRC" "$NGINX_CONF_DST"
systemctl restart nginx || true

#------------------------------------------------------------------------||
#-|| [6/8] Verify running containers
#------------------------------------------------------------------------||

echo "[6/8] Checking running containers..."
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

#------------------------------------------------------------------------||
#-|| [7/8] Show latest logs for key services
#------------------------------------------------------------------------||

echo "[7/8] Tailing logs for first 10 seconds..."
timeout 10 docker compose -f "$COMPOSE_FILE" logs --tail=20 || true

#------------------------------------------------------------------------||
#-|| [8/8] Cleanup old images
#------------------------------------------------------------------------||

echo "[8/8] Cleaning up unused Docker images..."
docker image prune -f || true

echo "====================================================================="
echo " Deployment completed successfully! "
echo "====================================================================="
