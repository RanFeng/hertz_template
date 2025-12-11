#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
export RUN_ENV=prod
BinaryName=hertz_demo
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}