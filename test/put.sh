#!/bin/bash

curl -v 127.0.0.1:12345/objects/test2 -XPUT -d"this is object v2"

curl 127.0.0.1:12345/locate/test2
echo
curl 127.0.0.1:12345/objects/test2
echo