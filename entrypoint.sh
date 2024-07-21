#!/bin/sh

/server > server.log 2>&1 &

nginx -g "daemon off;"