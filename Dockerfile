FROM mysql:8

WORKDIR /sql

COPY *.sql .

ENV MYSQL_ROOT_PASSWORD=password

VOLUME /var/lib/mysql

EXPOSE 3306
