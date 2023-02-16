#!/bin/sh

docker ps
echo $DOCKER_HOST
echo $COMPOSE_FILE
echo $REPO_ADDRESS
echo $EXECUTION_PATH
# "up -d" or "down"
echo $ACTION 

# TODO: ADD support to branch
# I may need the user to clone it
# use single branch to save space
git clone https://$GITHUB_TOKEN@$REPO_ADDRESS app
echo "repo ok: $?"
# cd $EXECUTION_PATH
cd app
docker compose ls --format json > /var/tmp/before.json
echo "collected before state: $?"
docker compose -f $COMPOSE_FILE $ACTION
echo "compose run as expected:  $?"
sleep 3
#sleep 60
cat /var/tmp/before.json
docker compose ls --format json > /var/tmp/after.json
echo "collected after state: $?"