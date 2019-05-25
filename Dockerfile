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
    fsync                           = 'off'             \n\
    synchronous_commit              = 'off'             \n\
    huge_pages                      = 'off'             \n\
    autovacuum                      = 'off'             \n\
    logging_collector               = 'off'             \n\
    archive_mode                    = 'off'             \n\
    ssl                             = 'off'             \n\
    row_security                    = 'off'             \n\
    parallel_leader_participation   = 'on'              \n\
    wal_compression                 = 'on'              \n\
    max_worker_processes            = '8'               \n\
    max_parallel_workers            = '8'               \n\
    max_wal_senders                 = '0'               \n\
    max_wal_senders                 = '0'               \n\
    effective_io_concurrency        = '0'               \n\
    wal_keep_segments               = '130'             \n\
    wal_level                       = 'minimal'         \n\
    log_min_messages                = 'panic'           \n\
    wal_writer_delay                = '2000ms'          \n\
    bgwriter_delay                  = '210ms'           \n\
    wal_buffers                     = '16MB'            \n\
    shared_buffers                  = '256 MB'          \n\
    effective_cache_size            = '1024 MB'         \n\
    work_mem                        = '32 MB'           \n\
    maintenance_work_mem            = '360 MB'          \n\
    " >> \
        "/etc/postgresql/11/main/postgresql.conf"


#jit                             = 'on'              \n\


EXPOSE 5432



VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

CMD service postgresql start && go run main.go