FROM postgres:16

COPY setup-db.sh /docker-entrypoint-initdb.d/01-setup-db.sh
COPY pg_schema /tmp/pg_schema/

RUN chmod +x /docker-entrypoint-initdb.d/01-setup-db.sh