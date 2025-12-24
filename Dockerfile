FROM postgres:16

COPY setup-db.sh /docker-entrypoint-initdb.d/
COPY pg_schema /docker-entrypoint-initdb.d/pg_schema/

RUN chmod +x /docker-entrypoint-initdb.d/setup-db.sh

