#!/bin/bash
set -e

echo "🚀 Init multiple databases..."

for file in /tmp/pg_schema/*.sql; do
  db_name=$(basename "$file" .sql)

  echo "📦 Creating database: $db_name"

  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE DATABASE "$db_name";
EOSQL

  echo "📥 Importing $file into $db_name"

  psql -v ON_ERROR_STOP=1 \
       --username "$POSTGRES_USER" \
       --dbname "$db_name" \
       -f "$file"

done

echo "✅ All databases initialized"