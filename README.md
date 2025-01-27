# Mono Monero

Mono Monero is a tor based online monero banking system.

## Default settings

Create a file called settings.json in the root of the project directory and add:

```
{
  "dsn": "file:/database.db",
  "port": 3000,
  "jwt_secret": "SECRETGOESHERE"
}

```
