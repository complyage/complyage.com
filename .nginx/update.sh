#!/bin/bash
set -e


echo ""
echo ""
echo ""
echo "====================================================================="
echo " ComplyAge Nginx Config Update Script "
echo "====================================================================="

REPO_DIR="/complyage/complyage.com"
NGINX_CONF_SRC="$REPO_DIR/.nginx/compose-proxy.conf"
NGINX_CONF_DST="/etc/nginx/conf.d/compose-proxy.conf"

#------------------------------------------------------------------------||
#-|| [1/3] Copy updated configuration
#------------------------------------------------------------------------||

echo "[1/3] Copying new Nginx configuration..."
if [ -f "$NGINX_CONF_SRC" ]; then
    cp -f "$NGINX_CONF_SRC" "$NGINX_CONF_DST"
    echo " → Copied: $NGINX_CONF_SRC → $NGINX_CONF_DST"
else
    echo " ❌ Source file not found: $NGINX_CONF_SRC"
    exit 1
fi

#------------------------------------------------------------------------||
#-|| [2/3] Test configuration syntax
#------------------------------------------------------------------------||

echo "[2/3] Testing Nginx configuration..."
echo "====================================================================="
nginx -t
echo "====================================================================="

#------------------------------------------------------------------------||
#-|| [3/3] Restart Nginx service
#------------------------------------------------------------------------||

echo "[3/3] Restarting Nginx..."
systemctl restart nginx

echo "✅ Nginx configuration successfully updated and reloaded!"
echo "====================================================================="
echo ""
echo ""
echo ""
