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
#-|| [3/8] Copy base Dockerfile into each service
#------------------------------------------------------------------------||

echo "[3/8] Copying base Dockerfile into service directories..."

BASE_FILE="$REPO_DIR/Dockerfile.base.service"

for service in api gate oauth .deployment; do
    TARGET_DIR="$REPO_DIR/$service"
    if [ -d "$TARGET_DIR" ]; then
        echo " → Copying to $service/"
        cp -f "$BASE_FILE" "$TARGET_DIR/Dockerfile"
    else
        echo " Skipping $service (directory not found)"
    fi
done

#------------------------------------------------------------------------||
#-|| [3.5/8] Clean up go.mod replace directives
#------------------------------------------------------------------------||

echo "[3.5/8] Cleaning local 'replace' directives from Go modules..."

for dir in api gate oauth deployment; do
   if [ -f "$REPO_DIR/$dir/go.mod" ]; then
      echo " → Cleaning $dir/go.mod"
      sed -i '/replace.*complyage.base/d' "$REPO_DIR/$dir/go.mod" || true
      sed -i '/replace.*ralphferrara.aria/d' "$REPO_DIR/$dir/go.mod" || true
      sed -i '/=> ../d' "$REPO_DIR/$dir/go.mod" || true
   fi
done

#------------------------------------------------------------------------||
#-|| [4/8] Clean and rebuild Docker images
#------------------------------------------------------------------------||

echo "[4/8] Rebuilding all containers..."
cd "$REPO_DIR"
docker compose -f "$COMPOSE_FILE" build --no-cache

#------------------------------------------------------------------------||
#-|| [5/8] Start Docker stack
#------------------------------------------------------------------------||

echo "[5/8] Starting Docker stack..."
docker compose -f "$COMPOSE_FILE" up -d

#------------------------------------------------------------------------||
#-|| [5.5/8] Build and deploy UI
#------------------------------------------------------------------------||

echo "[5.5/8] Building and deploying UI..."

UI_DIR="$REPO_DIR/ui"
UI_DIST_DIR="$UI_DIR/dist"
NGINX_UI_DIR="/var/www/complyage-ui"

if [ -d "$UI_DIR" ]; then
    echo " → Installing dependencies..."
    cd "$UI_DIR"
    npm ci --omit=dev || npm install --omit=dev

    echo " → Building production bundle..."
    npm run build

    echo " → Copying build output to Nginx web root..."
    mkdir -p "$NGINX_UI_DIR"
    rsync -av --delete "$UI_DIST_DIR/" "$NGINX_UI_DIR/"
else
    echo " ⚠️ UI directory not found at $UI_DIR — skipping build."
fi


#------------------------------------------------------------------------||
#-|| [6/8] Apply updated Nginx configuration
#------------------------------------------------------------------------||

echo "[6/8] Updating Nginx reverse proxy configuration..."
cp -f "$NGINX_CONF_SRC" "$NGINX_CONF_DST"
systemctl restart nginx || true

#------------------------------------------------------------------------||
#-|| [7/8] Verify running containers
#------------------------------------------------------------------------||

echo "[7/8] Checking running containers..."
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

#------------------------------------------------------------------------||
#-|| [8/8] Show logs & cleanup
#------------------------------------------------------------------------||

echo "[8/8] Tailing logs for first 10 seconds..."
timeout 10 docker compose -f "$COMPOSE_FILE" logs --tail=20 || true

echo "Cleaning up unused Docker images..."
docker image prune -f || true

echo "====================================================================="
echo " Deployment completed successfully! "
echo "====================================================================="
