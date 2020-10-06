FROM mysql:latest

COPY ./dtdump201002.sql /docker-entrypoint-initdb.d/dtdump201002.sql
#RUN /bin/bash -c "/usr/bin/mysqld_safe --skip-grant-tables" &&\
#    sleep 5s && \
#    mysql -u root -e "CREATE DATABASE dockertest" && \
#    mysql -u root dockertest < /docker-entrypoint-initdb.d/dtdump201002.sql