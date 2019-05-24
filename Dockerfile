FROM ubuntu:18.04

LABEL name="Martynov Anton"

ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

USER root
ENV DEBIAN_FRONTEND 'noninteractive'

RUN apt-get update -y
RUN apt-get install -y --no-install-recommends apt-utils

RUN apt-get install -y git wget

RUN wget https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.12.5.linux-amd64.tar.gz

ENV GOROOT /usr/local/go
ENV GOPATH /opt/go
ENV PATH $GOROOT/bin:$GOPATH/bin:/usr/local/go/bin:$PATH

ENV POSTGRESQLVERSION 11

RUN apt-get update && apt-get install -y wget gnupg &&     wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -

RUN echo "deb http://apt.postgresql.org/pub/repos/apt bionic-pgdg main" > /etc/apt/sources.list.d/PostgreSQL.list

RUN apt-get update && apt-get install -y postgresql-11

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
RUN printf "\n\
    fsync                = 'off'     \n\
    synchronous_commit   = 'off'     \n\
    autovacuum           = 'off'     \n\
    wal_level            = 'minimal' \n\
    wal_buffers          = '16MB'     \n\
    max_wal_senders      = '0'       \n\
    wal_writer_delay     = '2000ms'  \n\
    shared_buffers       = '512MB'   \n\
    effective_cache_size = '512MB'  \n\
    log_min_messages     = 'panic'   \n" >> \
        "/etc/postgresql/11/main/postgresql.conf"

EXPOSE 5432



VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

CMD service postgresql start && go run main.go