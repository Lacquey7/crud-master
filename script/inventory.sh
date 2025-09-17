#!/usr/bin/env bash
set -euo pipefail

export DEBIAN_FRONTEND=noninteractive

# ---- Variables ----
GO_VERSION="1.25.1"
GO_TAR="go${GO_VERSION}.linux-arm64.tar.gz"
GO_DOWNLOAD="https://go.dev/dl/${GO_TAR}"
INSTALL_DIR="/usr/local"
PROFILE_FILE="$HOME/.profile"
# -------------------

echo "[*] Update apt cache & base tools"
sudo apt-get update -y
sudo apt-get install -y curl tar ca-certificates build-essential git

echo "[*] Téléchargement de Go ${GO_VERSION} (ARM64)"
curl -fsSL "${GO_DOWNLOAD}" -o "/tmp/${GO_TAR}"

echo "[*] Suppression de l'ancienne installation (si existante)"
sudo rm -rf "${INSTALL_DIR}/go"

echo "[*] Installation dans ${INSTALL_DIR}/go"
sudo tar -C "${INSTALL_DIR}" -xzf "/tmp/${GO_TAR}"

# PATH runtime + persistance
if ! echo "$PATH" | grep -q "/usr/local/go/bin"; then
  export PATH="$PATH:/usr/local/go/bin"
fi
if ! grep -q "/usr/local/go/bin" "$PROFILE_FILE" 2>/dev/null; then
  echo 'export PATH="$PATH:/usr/local/go/bin"' >> "$PROFILE_FILE"
fi
if [ ! -f /etc/profile.d/go.sh ] || ! grep -q "/usr/local/go/bin" /etc/profile.d/go.sh; then
  echo 'export PATH="$PATH:/usr/local/go/bin"' | sudo tee /etc/profile.d/go.sh >/dev/null
fi

echo "[*] Vérification de l’installation Go"
which go || true
go version

# ---------------- Node.js + PM2 ----------------
if ! command -v node >/dev/null 2>&1; then
  echo "[*] Installation de Node.js (LTS via NodeSource)"
  curl -fsSL https://deb.nodesource.com/setup_lts.x | sudo -E bash -
  sudo apt-get install -y nodejs
else
  echo "[=] Node déjà installé: $(node -v)"
fi

echo "[*] Vérification de l’installation Node"
node -v
npm -v

if ! command -v pm2 >/dev/null 2>&1; then
  echo "[*] Installation de PM2 (global)"
  sudo npm install -g pm2
else
  echo "[=] PM2 déjà installé: $(pm2 -v)"
fi
pm2 -v

# ---------------- Docker ----------------
echo "[*] Dépôt Docker"
# (réinstall idempotente; pas de -y ici car ce n'est pas apt)
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list >/dev/null

sudo apt-get update -y
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

# Activer Docker et autoriser l'utilisateur vagrant
sudo systemctl enable --now docker
sudo usermod -aG docker vagrant || true

# Lancer PostgreSQL en arrière-plan, idempotent
if ! docker ps -a --format '{{.Names}}' | grep -q '^postgres$'; then
  docker run -d \
    --name postgres \
    -e POSTGRES_USER=mohamed \
    -e POSTGRES_PASSWORD=Azerty12@ \
    -e POSTGRES_DB=postgres \
    -p 5432:5432 \
    --restart unless-stopped \
    postgres
else
  docker start postgres || true
fi


# ---------------- Build & PM2 pour inventory-app ----------------
APP_DIR="/home/vagrant/inventory_app"
if [ -d "$APP_DIR" ]; then
  echo "[*] Build de l'application Go"
  cd "$APP_DIR"
  mkdir -p tmp
  go build -o ./tmp/main .

  echo "[*] Démarrage via PM2"
  if [ -f ecosystem.config.js ]; then
    # Si l'app existe déjà dans PM2, on relance proprement
    if pm2 list | grep -q "inventory-app"; then
      pm2 restart ecosystem.config.js --update-env || true
    else
      pm2 start ecosystem.config.js --update-env
    fi
  else
    # Fallback: lancer le binaire directement si pas de config PM2
    pm2 start ./tmp/main --name billing-app
  fi
  pm2 save
  # Activer le démarrage automatique au boot
  pm2 startup systemd -u vagrant --hp /home/vagrant >/dev/null 2>&1 || true
else
  echo "[!] Répertoire $APP_DIR introuvable, skip build/PM2"
fi

# Nettoyage
rm -f "/tmp/${GO_TAR}"

echo "[✔] Provisioning terminé avec succès"