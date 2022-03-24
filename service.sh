#!/bin/bash

CURRENTDATE=`date +"%D %T"`

# Getting absolute path
SOURCE=${BASH_SOURCE[0]}
while [ -h "$SOURCE" ]; do
  TARGET=$(readlink "$SOURCE")
  if [[ $TARGET == /* ]]; then
    SOURCE=$TARGET
  else
    DIR=$( dirname "$SOURCE" )
    SOURCE=$DIR/$TARGET
  fi
done
RDIR=$( dirname "$SOURCE" )
DIR=$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )

# Install dependencies
if [ "$1" == "install" ]
then
    echo [${CURRENTDATE}] - Installation started...
    apt-get install docker
    apt-get install golang
    go mod download
    echo [${CURRENTDATE}] - Installation completed
fi

# Build postgresql docker service
if [ "$1" == "build_db" ]
then
    echo [${CURRENTDATE}] - Database building started...
    systemctl stop postgresql
    docker run --name go_example_db -e POSTGRES_PASSWORD=secret -d -p 5432:5432 -v /volumes/test/postgresql/data:/var/lib/postgresql/data  postgres
    echo [${CURRENTDATE}] - Database build completed
fi

# Run postgresql service after stoping local postgresql service
if [ "$1" == "run_db" ]
then
    systemctl stop postgresql
    docker start go_example_db
    echo [${CURRENTDATE}] - Database started
fi

# Build main go service
if [ "$1" == "build" ]
then
    echo [${CURRENTDATE}] - Build started
    cd cmd/core
    go build -o $DIR/bin/example-microservice
    echo [${CURRENTDATE}] - Build completed
fi

# Run go microservice
if [ "$1" == "run" ]
then
    echo [${CURRENTDATE}] - Running service...
    ./bin/example-microservice
fi