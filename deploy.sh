#!/bin/bash
docker stop time || true
docker rm `docker ps --no-trunc -aq`  || true

docker rmi -f time-image || true
docker build -t time-image /home/ec2-user/.
docker run -p 80:8080 --name=time -d time-image