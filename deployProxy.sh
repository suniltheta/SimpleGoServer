#!/bin/bash
docker stop proxy || true
docker rm `docker ps --no-trunc -aq`  || true

docker rmi -f proxy-image || true
docker build -t proxy-image /home/ec2-user/.
docker run -p 80:8080 -p 443:8081 --name=proxy -d proxy-image