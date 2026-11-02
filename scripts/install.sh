#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
BIN_DIR="${PGP_INSTALL_DIR:-$HOME/.local/bin}"
NAME=pgp

mkdir -p "$BIN_DIR"
(cd "$ROOT" && go build -ldflags="-s -w" -o "$BIN_DIR/$NAME" .)

echo "Installed $BIN_DIR/$NAME"

PLUGIN_SRC="$ROOT/oh-my-zsh/pgp.plugin.zsh"
PLUGIN_DST="$HOME/.oh-my-zsh/custom/plugins/pgp/pgp.plugin.zsh"
if [[ -d "$HOME/.oh-my-zsh/custom/plugins" ]]; then
  mkdir -p "$(dirname "$PLUGIN_DST")"
  ln -sf "$PLUGIN_SRC" "$PLUGIN_DST"
  echo "Linked Oh My Zsh plugin -> $PLUGIN_DST"
  if ! grep -q 'plugins=(.*pgp' "$HOME/.zshrc" 2>/dev/null; then
    echo ""
    echo "Add pgp to your plugins in ~/.zshrc, e.g.:"
    echo "  plugins=(git pgp ...)"
  fi
fi

if ! command -v "$NAME" >/dev/null 2>&1; then
  echo "Ensure $BIN_DIR is on PATH (e.g. export PATH=\"\$HOME/.local/bin:\$PATH\" in ~/.zshrc)"
fi
