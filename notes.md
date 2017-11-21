### Steps to deploy the application

Things to consider<br>
AMI is created which has docker image of golang. So while running Ansible scripts git clone this repo and copy Dockerfile into destination server. Copy `/src/time` into `app` folder of detination server. Then run following docker commands as stopping running instance, deleting it, building docker image and running a container out of this new image.<br><br>

deploy.sh
```
#!/bin/bash
docker stop time || true
docker rm `docker ps --no-trunc -aq`  || true

docker rmi -f time-image || true
docker build -t time-image /home/ec2-user/.
docker run -p 80:8080 --name=time -d time-image
```

Note: Docker command to attach to a terminal of a running container <br>
`$  docker exec -i -t <name> /bin/bash`


### Run ansible playbook after git cloning this repo.
Make sure before runnig ansible playbook you have inventory of servers<br>
To deploy Time server<br>
`$ sudo ansible-playbook SimpleGoServer/deploy.yml -i inventory`
To deploy Proxy server<br>
`$ sudo ansible-playbook SimpleGoServer/deploy-proxy.yml -i inventory`
<br><br><br>
deploy.yml
```
---
- hosts: prod
  tasks:
  - name: Upload src app
    copy: src=src/time dest=/home/ec2-user/app

  - name: Upload Dockerfile
    copy: src=Dockerfile dest=/home/ec2-user/Dockerfile mode=0755

  - name: Deploy Docker container
    become_user: ec2-user
    script: deploy.sh
```