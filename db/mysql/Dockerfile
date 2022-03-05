FROM mysql

#set enviroment
ENV MYSQL_ROOT_PASSWORD admin

RUN mkdir /db/sql

ADD ["db/mysql/blog-service.sql","/usr/db/sql/blog-service.sql"]

RUN /etc/init.d/mysql start
RUN mysql -u root