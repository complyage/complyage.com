#!/bin/bash
set -e

echo "====================================================================="
echo " ComplyAge Deployment Script "
echo "====================================================================="

REPO_DIR="/complyage/complyage.com"
NGINX_CONF_SRC="$REPO_DIR/.nginx/compose-proxy.conf"
NGINX_CONF_DST="/etc/nginx/conf.d/compose-proxy.conf"

#------------------------------------------------------------------------||
#-|| [1/14] Shutdown the /complyage/complyage.com/docker-compose.yaml
#------------------------------------------------------------------------||

echo "[1/14] Shutting down Docker stack..."
docker compose -f "$REPO_DIR/docker-compose.yml" down || true

#------------------------------------------------------------------------||
#-|| [2/14] Rebuilding Docker stack
#------------------------------------------------------------------------||

echo "[2/14] Rebuilding Docker stack..."
docker compose -f "$REPO_DIR/docker-compose.yml" build --no-cache

#------------------------------------------------------------------------||
#-|| [3/14] Starting Docker stack
#------------------------------------------------------------------------||

echo "[3/14] Starting Docker stack..."
docker compose -f "$REPO_DIR/docker-compose.yml" up -d

#------------------------------------------------------------------------||
#-|| [4/14] Stopping Nginx
#------------------------------------------------------------------------||

echo "[4/14] Stopping Nginx..."
systemctl stop nginx || true

#------------------------------------------------------------------------||
#-|| [5/14] Copying new Nginx proxy configuration
#------------------------------------------------------------------------||

echo "[5/14] Copying proxy configuration..."
cp -f "$NGINX_CONF_SRC" "$NGINX_CONF_DST"

#------------------------------------------------------------------------||
#-|| [6/14] Building Go services (.deployment)
#------------------------------------------------------------------------||

echo "[6/14] Building .deployment service..."
cd "$REPO_DIR/.deployment"
go build ./...
echo "[7/14] Running .deployment service..."
nohup go run ./... > /var/log/deploy.log 2>&1 &

#------------------------------------------------------------------------||
#-|| [8/14] Building API service
#------------------------------------------------------------------------||

echo "[8/14] Building API service..."
cd "$REPO_DIR/api"
go build ./...
echo "[9/14] Running API service..."
nohup go run ./... > /var/log/api.log 2>&1 &

#------------------------------------------------------------------------||
#-|| [10/14] Building Gate service
#------------------------------------------------------------------------||

echo "[10/14] Building Gate service..."
cd "$REPO_DIR/gate"
go build ./...
echo "[11/14] Running Gate service..."
nohup go run ./... > /var/log/gate.log 2>&1 &

#------------------------------------------------------------------------||
#-|| [12/14] Building OAuth service
#------------------------------------------------------------------------||

echo "[12/14] Building OAuth service..."
cd "$REPO_DIR/oauth"
go build ./...
echo "[13/14] Running OAuth service..."
nohup go run ./... > /var/log/oauth.log 2>&1 &

#------------------------------------------------------------------------||
#-|| [14/14] Restarting Nginx and cleanup
#------------------------------------------------------------------------||

echo "[14/14] Restarting Nginx..."
systemctl start nginx || true

echo "Cleaning dangling Docker images..."
docker image prune -f || true

echo "====================================================================="
echo " Deployment completed successfully! "
echo "====================================================================="
