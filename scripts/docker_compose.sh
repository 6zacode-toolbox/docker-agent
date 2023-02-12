#!/bin/sh

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
#git clone https://[TOKEN]@github.com/[REPO-OWNER]/[REPO-NAME] app
# cd $EXECUTION_PATH
# docker compose -f $COMPOSE_FILE $ACTION
	

docker compose ls --format json > /var/tmp/after.json