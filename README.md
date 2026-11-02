# pgp (go-postgres-parser)

Local Go CLI that **parses** a Postgres connection URL and prints the parts. **No network, no database connection.**

## Install

```bash
./scripts/install.sh
```

## Usage

```bash
pgp 'postgres://neondb_owner:YOUR_PASSWORD@ep-....neon.tech:5432/neondb?sslmode=require'
```

Output:

```
Host: ep-....neon.tech
Port: 5432
Database Name: neondb
Username: neondb_owner
Password: YOUR_PASSWORD
```

Quote the URL in zsh so special characters in the password are preserved.
