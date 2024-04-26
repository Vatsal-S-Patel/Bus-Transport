FROM postgres:14.11

WORKDIR /docker-entrypoint-initdb.d

COPY createBacancy.sql  .

COPY dump.sql .

EXPOSE 5432
