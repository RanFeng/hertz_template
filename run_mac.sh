#!/bin/bash
RUN_NAME=hertz_demo
mkdir -p output/bin output/conf
cp script/* output/
cp conf/* output/conf/
chmod +x output/bootstrap.sh
go build -o output/bin/${RUN_NAME}

./output/bootstrap_local.sh