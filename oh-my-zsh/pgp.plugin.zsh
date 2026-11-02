# pgp — parse postgres URLs locally (no network)

if (( $+commands[pgp] )); then
  _pgp() {
    _message 'postgres URL'
  }
  compdef _pgp pgp 2>/dev/null || true
fi
