FROM ubuntu:18.04

LABEL name="Martynov Anton"

ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

USER root
ENV DEBIAN_FRONTEND 'noninteractive'

RUN apt-get update -y
RUN apt-get install -y --no-install-recommends apt-utils

RUN apt-get install -y wget
RUN apt-get install -y git

RUN wget https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.12.5.linux-amd64.tar.gz

ENV GOROOT /usr/local/go
ENV GOPATH /opt/go
ENV PATH $GOROOT/bin:$GOPATH/bin:/usr/local/go/bin:$PATH

ENV POSTGRESQLVERSION 10
RUN apt-get install -y postgresql-$POSTGRESQLVERSION

WORKDIR /tech-db-forum
COPY . .

EXPOSE 5000

RUN go get -u

USER postgres
RUN /etc/init.d/postgresql start &&\
    psql --command "CREATE USER sayonara WITH SUPERUSER PASSWORD 'boy';" &&\
    createdb -O sayonara techno &&\
    psql --dbname=techno --echo-all --command 'create extension if not exists "citext";' &&\
    /etc/init.d/postgresql stop

USER root
RUN echo "host all all 0.0.0.0/0 md5" >> /etc/postgresql/$POSTGRESQLVERSION/main/pg_hba.conf &&\
    echo "listen_addresses='*'" >> /etc/postgresql/$POSTGRESQLVERSION/main/postgresql.conf &&\
    echo "shared_buffers=256MB" >> /etc/postgresql/$POSTGRESQLVERSION/main/postgresql.conf &&\
    echo "fsync=off" >> /etc/postgresql/$POSTGRESQLVERSION/main/postgresql.conf &&\
    echo "synchronous_commit=off" >> /etc/postgresql/$POSTGRESQLVERSION/main/postgresql.conf &&\
    echo "full_page_writes=off" >> /etc/postgresql/$POSTGRESQLVERSION/main/postgresql.conf &&\
    echo "unix_socket_directories = '/var/run/postgresql'" >> /etc/postgresql/$POSTGRESQLVERSION/main/postgresql.conf
EXPOSE 5432

VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

CMD service postgresql start && go run main.go