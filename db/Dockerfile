FROM mysql:8.0

ADD ./my.cnf /etc/mysql/conf.d/my.cnf

CMD ["mysqld"]

RUN ./docker-entrypoint-initdb.d/init.sh
