# start with base image
FROM mysql

# import data into container
# All scripts in docker-entrypoint-initdb.d/ are automatically executed during container startup
COPY ./schema/base_schema.sql /docker-entrypoint-initdb.d/
