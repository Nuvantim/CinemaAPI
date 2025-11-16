#!/bin/bash
set -e
set -u

# Function to create a PostgreSQL user and database, and optionally import SQL data
function create_user_and_database() {
  local database=$1
  echo "🚀 Creating user and database: '$database'"

  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "postgres" <<-EOSQL
    CREATE USER $database WITH PASSWORD '$POSTGRES_PASSWORD';
    CREATE DATABASE $database;
    GRANT ALL PRIVILEGES ON DATABASE $database TO $database;
EOSQL

  local sql_file="/docker-entrypoint-initdb.d/pg_schema/${database}.sql"

  # Check if the SQL file exists and is not empty
  if [ -s "$sql_file" ]; then
    echo "📦 Importing data from $sql_file into '$database'"
    psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname="$database" -f "$sql_file"
  else
    # If the file is empty or missing, create a basic table as a placeholder
    echo "⚙️  No SQL file found or file is empty. Generating default table for '$database'"
    psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname="$database" <<-EOSQL
      CREATE TABLE ${database} (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
      );
EOSQL
  fi
}

echo "🔍 Scanning for SQL files in /docker-entrypoint-initdb.d/pg_schema/ ..."
for sql_file in /docker-entrypoint-initdb.d/pg_schema/*.sql; do
  db_name=$(basename "$sql_file" .sql)
  create_user_and_database "$db_name"
done

echo "✅ All databases and tables have been successfully created!"

