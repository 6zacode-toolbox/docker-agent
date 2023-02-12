#!/bin/sh

docker ps
docker compose ls --format json > /var/tmp/before.json
echo $DOCKER_HOST
echo $COMPOSE_FILE
echo $REPO_ADDRESS
echo $EXECUTION_PATH
echo $GITHUB_TOKEN
echo $CRD_API_VERSION
echo $CRD_NAMESPACE
echo $CRD_NAME
echo $CRD_RESOURCE
# "up -d" or "down"
echo $ACTION 

# TODO: ADD support to branch
# I may need the user to clone it
# use single branch to save space
git clone https://$GITHUB_TOKEN@$REPO_ADDRESS app
# cd $EXECUTION_PATH
cd app
docker compose -f $COMPOSE_FILE $ACTION
sleep 10
docker ps	
sleep 60
cat /var/tmp/before.json
#docker compose ls --format json > /var/tmp/after.json