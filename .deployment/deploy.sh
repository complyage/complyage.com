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
#-|| [1/9] Stop running containers
#------------------------------------------------------------------------||

echo "[1/9] Stopping any existing Docker containers..."
docker compose -f "$COMPOSE_FILE" down || true

#------------------------------------------------------------------------||
#-|| [2/9] Update or clone Aria Framework
#------------------------------------------------------------------------||

echo "[2/9] Updating Aria Framework..."
if [ -d "$ARIA_DIR/.git" ]; then
    echo " → Aria already exists — pulling latest changes..."
    cd "$ARIA_DIR" && git reset --hard && git pull origin main
else
    echo " → Aria not found — cloning fresh..."
    git clone https://github.com/ralphferrara/aria.git "$ARIA_DIR"
fi

#------------------------------------------------------------------------||
#-|| [3/9] Copy base Dockerfile into each service
#------------------------------------------------------------------------||

echo "[3/9] Copying base Dockerfile into service directories..."

BASE_FILE="$REPO_DIR/Dockerfile.base.service"

for service in api gate oauth .deployment; do
    TARGET_DIR="$REPO_DIR/$service"
    if [ -d "$TARGET_DIR" ]; then
        echo " → Copying to $service/"
        cp -f "$BASE_FILE" "$TARGET_DIR/Dockerfile"
    else
        echo " ⚠️ Skipping $service (directory not found)"
    fi
done

#------------------------------------------------------------------------||
#-|| [4/9] Clean up go.mod replace directives
#------------------------------------------------------------------------||

echo "[4/9] Cleaning local 'replace' directives from Go modules..."

for dir in api gate oauth deployment; do
   if [ -f "$REPO_DIR/$dir/go.mod" ]; then
      echo " → Cleaning $dir/go.mod"
      sed -i '/replace.*complyage.base/d' "$REPO_DIR/$dir/go.mod" || true
      sed -i '/replace.*ralphferrara.aria/d' "$REPO_DIR/$dir/go.mod" || true
      sed -i '/=> ../d' "$REPO_DIR/$dir/go.mod" || true
   fi
done

#------------------------------------------------------------------------||
#-|| [4.5/9] Copy shared configs and secure credentials into each service directory
#------------------------------------------------------------------------||

echo "[4.5/9] Copying shared .env, config.json, and .secure directory into each service..."

CONFIG_SRC="/complyage/.config"
SECURE_SRC="/complyage/.secure"
SERVICES=("api" "gate" "oauth" ".deployment")

for service in "${SERVICES[@]}"; do
    TARGET_DIR="/complyage/complyage.com/$service"
    if [ -d "$TARGET_DIR" ]; then
        echo " → Processing $service..."

        # Copy .env
        if [ -f "$CONFIG_SRC/.env" ]; then
            cp -f "$CONFIG_SRC/.env" "$TARGET_DIR/.env"
            echo "   ✓ Copied .env"
        else
            echo "   ⚠️ Missing $CONFIG_SRC/.env"
        fi

        # Copy config.json
        if [ -f "$CONFIG_SRC/config.json" ]; then
            cp -f "$CONFIG_SRC/config.json" "$TARGET_DIR/config.json"
            echo "   ✓ Copied config.json"
        else
            echo "   ⚠️ Missing $CONFIG_SRC/config.json"
        fi

        # Copy .secure directory
        if [ -d "$SECURE_SRC" ]; then
            mkdir -p "$TARGET_DIR/.secure"
            cp -r "$SECURE_SRC/"* "$TARGET_DIR/.secure/" 2>/dev/null || true
            echo "   ✓ Copied .secure directory"
        else
            echo "   ⚠️ Missing $SECURE_SRC directory"
        fi

    else
        echo " ⚠️ Skipping $service (directory not found)"
    fi
done


#------------------------------------------------------------------------||
#-|| [5/9] Clean and rebuild Docker images
#------------------------------------------------------------------------||

echo "[5/9] Rebuilding all containers..."
cd "$REPO_DIR"
docker compose -f "$COMPOSE_FILE" build --no-cache

#------------------------------------------------------------------------||
#-|| [6/9] Start Docker stack
#------------------------------------------------------------------------||

echo "[6/9] Starting Docker stack..."
docker compose -f "$COMPOSE_FILE" up -d

#------------------------------------------------------------------------||
#-|| [7/9] Build UI (npm install + build)
#------------------------------------------------------------------------||

echo "[7/9] Building UI..."

UI_DIR="$REPO_DIR/ui"

if [ -d "$UI_DIR" ]; then
    echo " → Installing dependencies..."
    cd "$UI_DIR"

    # Always use npm ci for deterministic installs
    # Include devDependencies so Vite is available
    if [ -f "package-lock.json" ]; then
        npm ci
    else
        npm install
    fi

    # Build using Vite
    echo " → Building production bundle..."
    npx vite build

    echo " ✅ UI build complete: $UI_DIR/dist"
else
    echo " ⚠️ UI directory not found at $UI_DIR — skipping build."
fi

#------------------------------------------------------------------------||
#-|| [7.5/9] Build Documentation (Vite)
#------------------------------------------------------------------------||

echo "[7.5/9] Building Documentation..."

DOCS_DIR="$REPO_DIR/documentation"

if [ -d "$DOCS_DIR" ]; then
    echo " → Installing dependencies for documentation..."
    cd "$DOCS_DIR"

    if [ -f "package-lock.json" ]; then
        npm ci
    else
        npm install
    fi

    echo " → Building production bundle..."
    npx vite build

    echo " ✅ Documentation build complete: $DOCS_DIR/dist"
else
    echo " ⚠️ Documentation directory not found — skipping."
fi

#------------------------------------------------------------------------||
#-|| [8/9] Apply updated Nginx configuration
#------------------------------------------------------------------------||

echo "[8/9] Updating Nginx reverse proxy configuration..."
cp -f "$NGINX_CONF_SRC" "$NGINX_CONF_DST"
systemctl restart nginx || true

#------------------------------------------------------------------------||
#-|| [9/9] Verify containers, show logs, and cleanup
#------------------------------------------------------------------------||

echo "[9/9] Checking running containers..."
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

echo "Tailing logs for first 10 seconds..."
timeout 10 docker compose -f "$COMPOSE_FILE" logs --tail=20 || true

echo "Cleaning up unused Docker images..."
docker image prune -f || true

echo "====================================================================="
echo " Deployment completed successfully! "
echo "====================================================================="
